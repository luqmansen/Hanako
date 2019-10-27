package models_mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

type AnimeDAO struct {
	Server string
	Database string
	Password string
	Username string
}

var db *mgo.Database
var dao = AnimeDAO{}


const COLLECTION  = "anime"

func init(){
	dao.Server = os.Getenv("mongo_server")
	dao.Database = os.Getenv("mongo_database")
	dao.Username = os.Getenv("mongo_username")
	dao.Password = os.Getenv("mongo_password")
	dao.Connect()
}

func (a *AnimeDAO) Connect(){
	session, err := mgo.Dial(a.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(a.Database)
}

func(a *AnimeDAO) FindAll() ([]Anime, error){
	var animes [] Anime
	err := db.C(COLLECTION).Find(bson.M{}).Limit(20).All(&animes)
	return animes,err
}