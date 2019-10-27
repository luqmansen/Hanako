package models_mongo

import "gopkg.in/mgo.v2/bson"

type Anime struct {
	Id bson.ObjectId `bson:"_id" json:"id"`
	Sources []string `bson:"sources" json:"sources"`
	Type string `bson:"type" json:"type"`
	Title string `bson:"title" json:"title"`
	Picture string `bson:"picture" json:"picture"`
	Relations []string `bson:"relations" json:"relations"`
	Thumbnail string `bson:"thumbnail" json:"thumbnail"`
	Episodes int `bson:"episodes" json:"episodes"`
	Synonyms []string `bson:"synonyms" json:"synonyms"`
}

