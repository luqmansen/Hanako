package main

import (
	"anime/handler"
	anime "anime/proto/anime"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	service := micro.NewService(
		micro.Name("hanako.srv.anime"),
	)

	service.Init()

	err := anime.RegisterAnimeServiceHandler(service.Server(), new(handler.Anime))
	if err != nil {
		log.Error(err)
	}

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
