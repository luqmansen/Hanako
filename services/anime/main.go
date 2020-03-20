package main

import (
	"github.com/joho/godotenv"
	"github.com/luqmansen/hanako/services/anime/handler"
	"github.com/luqmansen/hanako/services/anime/handler/datastore"
	proto "github.com/luqmansen/hanako/services/anime/proto"
	"github.com/micro/go-micro/v2"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	if e := godotenv.Load(); e != nil {
		logrus.Panicf("Couldn't load env %s", e.Error())
	}

}

func main() {
	srv := micro.NewService(
		micro.Name("hanako.srv.anime"),
	)

	session, err := datastore.CreateSession()
	defer session.Close()
	if err != nil {
		logrus.Errorf("Error connecting to datastore: %v", err.Error())
	}
	logrus.Infof("Successfully connected to database: %s", os.Getenv("mongo_dbname"))

	srv.Init()

	err = proto.RegisterAnimeServiceHandler(srv.Server(), &handler.AnimeService{session})
	if err != nil {
		logrus.Errorf("Error registering handler: %s", err.Error())
	}

	if err := srv.Run(); err != nil {
		logrus.Fatalf("Error running service: %s", err.Error())
	}
}
