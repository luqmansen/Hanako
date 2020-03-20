package datastore

import (
	proto "github.com/luqmansen/hanako/services/anime/proto"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
)

const (
	CollectionName = "anime"
)

var DbName = os.Getenv("mongo_dbname")

type AnimeRepository struct {
	Session *mgo.Session
}

type Repository interface {
	Create(*proto.Anime) error
	FindAll() ([]*proto.Anime, error)
	FindByQuery(request map[string]interface{}) ([]*proto.Anime, error)
	Close()
}

func (repo *AnimeRepository) collection() *mgo.Collection {
	return repo.Session.DB(DbName).C(CollectionName)
}

//Create new Anime
func (repo *AnimeRepository) Create(anime *proto.Anime) error {
	return repo.collection().Insert(anime)
}

//Close Session
func (repo *AnimeRepository) Close() {
	repo.Session.Close()
}

//FindAll anime possible with limit 20
func (repo *AnimeRepository) FindAll() ([]*proto.Anime, error) {
	var results []*proto.Anime

	err := repo.collection().Find(bson.M{}).Limit(20).All(&results)
	if err != nil {
		logrus.Error(err.Error())
	}

	return results, err
}

//FindByQuery get anime by query provided
func (repo *AnimeRepository) FindByQuery(request map[string]interface{}) ([]*proto.Anime, error) {
	var results []*proto.Anime

	query := queryFormater(request)

	err := repo.collection().Find(query).All(&results)
	if err != nil {
		logrus.Error(err.Error())
	}

	return results, err
}
