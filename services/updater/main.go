package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"updater/handler"
	"updater/subscriber"

	updater "updater/proto/updater"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("hanako.srv.updater"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	updater.RegisterUpdaterHandler(service.Server(), new(handler.Updater))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("hanako.srv.updater", service.Server(), new(subscriber.Updater))

	// Register Function as Subscriber
	micro.RegisterSubscriber("hanako.srv.updater", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
