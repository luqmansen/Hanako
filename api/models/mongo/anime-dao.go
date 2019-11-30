package mongo

import (
	"fmt"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
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
	fmt.Println(db)

	index := mgo.Index{
		Key:  []string{"$text:title"},
		Name: "titleIndex",
	}

	err = db.C(COLLECTION).EnsureIndex(index)
	if err != nil {
		//panic(err)
	}

}

func (a *AnimeDAO) FindAll() ([]Anime, error) {
	var results []Anime
	err := db.C(COLLECTION).Find(bson.M{}).Limit(20).All(&results)
	if err != nil {
		fmt.Println(err)
	}
	if results == nil{
		return []Anime{}, nil
	}
	return results, err
}

func (a *AnimeDAO) FindByQuery(queries map[string]interface{}) ([]Anime, error) {

	var results []Anime
	var query bson.M
	if queries["title"] != "" && queries["type"] != "" {
		query = bson.M{
			"$and": []interface{}{
				bson.M{
					"$text": bson.M{"$search": queries["title"]},
				},
				bson.M{"type": queries["type"]},
			},
		}
	} else if queries["title"] != "" {
		query = bson.M{
			"$text": bson.M{
				"$search": queries["title"],
			},
		}
	} else {
		query = bson.M{
			"type": queries["type"],
		}
	}

	fmt.Println(query)

	err := db.C(COLLECTION).Find(query).All(&results)
	if err != nil {
		fmt.Println(err)
	}
	return results, err
}
