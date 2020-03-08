package datastore

import (
	anime "anime/proto/anime"
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"unicode"
)

type AnimeDAO struct {
	Server   string
	Database string
	Password string
	Username string
}

var db *mgo.Database
var dao = AnimeDAO{}

const COLLECTION = "anime"

func init() {
	e := godotenv.Load()
	if e != nil {
		panic(e)
	}

	dao.Connect()
}

func (a *AnimeDAO) Connect() {

	session, err := mgo.Dial(os.Getenv("mongo_url"))
	if err != nil {
		panic(err)
	}
	db = session.DB(os.Getenv("mongo_dbname"))

	index := mgo.Index{
		Key:  []string{"$text:title"},
		Name: "titleIndex",
	}

	err = db.C(COLLECTION).EnsureIndex(index)
	if err != nil {
		log.Fatal(err)
	}

}

func (a *AnimeDAO) FindAll() ([]*anime.Anime, error) {
	var results []*anime.Anime
	err := db.C(COLLECTION).Find(bson.M{}).Limit(20).All(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results, err
}

func (a *AnimeDAO) FindByQuery(request map[string]interface{}) ([]*anime.Anime, error) {

	var results []*anime.Anime

	query := queryFormater(request)

	err := db.C(COLLECTION).Find(query).All(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results, err
}

func queryFormater(request map[string]interface{}) bson.M {
	var query bson.M

	if request["title"] != "" && request["type"] != "" {
		query = bson.M{
			"$and": []interface{}{
				bson.M{
					"$text": bson.M{"$search": request["title"]},
				},
				bson.M{"type": request["type"]},
			},
		}
	} else if request["title"] != "" {
		// Format string if it is contain more than one word to be queried as title
		s := fmt.Sprintf("%v", request["title"])
		for _, v := range s {
			if unicode.IsSpace(v) {
				s = "\"" + s + "\""
				break
			}
		}
		query = bson.M{
			"$text": bson.M{
				"$search": s,
			},
		}
	} else {
		query = bson.M{
			"type": request["type"],
		}
	}
	return query
}
