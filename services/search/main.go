package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	"search/handler"
	"search/subscriber"

	search "search/proto/search"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("hanako.srv.search"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	search.RegisterSearchHandler(service.Server(), new(handler.Search))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("hanako.srv.search", service.Server(), new(subscriber.Search))

	// Register Function as Subscriber
	micro.RegisterSubscriber("hanako.srv.search", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
