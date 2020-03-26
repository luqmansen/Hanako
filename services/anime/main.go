package main

import (
	"github.com/joho/godotenv"
	"github.com/luqmansen/hanako/services/anime/handler"
	"github.com/luqmansen/hanako/services/anime/handler/datastore"
	proto "github.com/luqmansen/hanako/services/anime/proto"
	"github.com/luqmansen/hanako/tracer"
	"github.com/micro/go-micro/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	servicename = "hanako.srv.anime"
)

func init() {
	if e := godotenv.Load(); e != nil {
		logrus.Panicf("Couldn't load env %s", e.Error())
	}

}

func main() {
	t, io, err := tracer.NewTracer(servicename, "localhost:6381")
	if err != nil {
		logrus.Fatal(err)
	}
	defer io.Close()

	opentracing.SetGlobalTracer(t)

	srv := micro.NewService(
		micro.Name(servicename),
	)

	session, err := datastore.CreateSession()
	if err != nil {
		logrus.Errorf("Error connecting to datastore: %v", err.Error())
	}
	defer session.Close()
	logrus.Infof("Successfully connected to database: %s", os.Getenv("mongo_dbname"))

	srv.Init()

	err = proto.RegisterAnimeServiceHandler(srv.Server(), &handler.AnimeService{Session: session})
	if err != nil {
		logrus.Errorf("Error registering handler: %s", err.Error())
	}

	if err := srv.Run(); err != nil {
		logrus.Fatalf("Error running service: %s", err.Error())
	}
}
