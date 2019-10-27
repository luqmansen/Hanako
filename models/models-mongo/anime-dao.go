package models_mongo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

type AnimeDAO struct {
	Server string
	Database string
}

var db *mgo.Database
var config = Config{}
var dao = AnimeDAO{}


const COLLECTION  = "anime"

func init(){
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
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