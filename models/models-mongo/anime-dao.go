package models_mongo

import (
	"fmt"
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


const COLLECTION  = "anime"

func init(){
	dao.Connect()
}

func (a *AnimeDAO) Connect(){

	session, err := mgo.Dial(os.Getenv("mongo_url"))
	if err != nil {
		fmt.Println(err)
	}
	db = session.DB(os.Getenv("mongo_dbname"))
}

func(a *AnimeDAO) FindAll() ([]Anime, error){
	var animes [] Anime
	err := db.C(COLLECTION).Find(bson.M{}).Limit(20).All(&animes)
	return animes,err
}