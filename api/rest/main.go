package main

import (
	"github.com/luqmansen/hanako/api/rest/handler"
	rest "github.com/luqmansen/hanako/api/rest/proto"
	"github.com/luqmansen/hanako/api/rest/utils"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/errors"
	"log"
)

func main() {
	service := micro.NewService(
		micro.Name(utils.ServiceName),
	)
	service.Init()
	if err := rest.RegisterRestHandler(service.Server(), &handler.Anime{Client: service.Client()}); err != nil {
		log.Fatal(errors.InternalServerError(utils.ServiceName, err.Error()))
	}

	if err := service.Run(); err != nil {
		log.Fatal(errors.InternalServerError(utils.ServiceName, err.Error()))
	}

}
