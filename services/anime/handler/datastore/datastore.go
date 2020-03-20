package datastore

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"os"
)

func init() {
	session, err := mgo.Dial(os.Getenv("mongo_url"))
	if err != nil {
		logrus.Panicf("Can't connect to database %s", err.Error())
	}
	index := mgo.Index{
		Key:  []string{"$text:title"},
		Name: "titleIndex",
	}
	err = session.DB("hanako").Login("root", "root")
	if err != nil {
		logrus.Panicf("Can't authenticate to database %s", err.Error())
	}
	err = session.DB(DbName).C(CollectionName).EnsureIndex(index)
	if err != nil {
		logrus.Panicf("Can't perform operation %s", err.Error())
	}
}

func CreateSession() (*mgo.Session, error) {
	session, err := mgo.Dial(os.Getenv("mongo_url"))
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
