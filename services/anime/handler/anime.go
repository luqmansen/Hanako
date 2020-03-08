package handler

import (
	models "anime/handler/datastore"
	anime "anime/proto/anime"
	"context"
	"errors"
	"github.com/micro/go-micro/metadata"
	"golang.org/x/net/trace"
)

type Anime struct {
	AnimeList []*anime.Anime
}

var dao = models.AnimeDAO{}

func (a *Anime) GetAll(ctx context.Context, req *anime.Request, resp *anime.Results) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}
	data, err := dao.FindAll()
	if err != nil {
		return errors.New("error when retrieving data " + err.Error())
	}
	resp.Animes = data
	return nil
}

func (a *Anime) GetAnimes(ctx context.Context, req *anime.Request, resp *anime.Results) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}
	query := map[string]interface{}{
		"title": req.GetTitle(),
		"type":  req.GetType(),
	}

	data, err := dao.FindByQuery(query)
	if err != nil {
		return errors.New("error when retrieving data " + err.Error())
	}
	resp.Animes = data
	return nil
}
