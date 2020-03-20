package handler

import (
	"context"
	"errors"
	"github.com/luqmansen/hanako/services/anime/handler/datastore"
	proto "github.com/luqmansen/hanako/services/anime/proto"
	"github.com/micro/go-micro/v2/metadata"
	"golang.org/x/net/trace"
	"gopkg.in/mgo.v2"
)

type AnimeService struct {
	Session *mgo.Session
}

func (s *AnimeService) GetRepo() datastore.Repository {
	return &datastore.AnimeRepository{Session: s.Session.Clone()}
}

func (s *AnimeService) GetAll(ctx context.Context, req *proto.Request, resp *proto.Results) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}

	repo := s.GetRepo()
	defer repo.Close()

	data, err := repo.FindAll()
	if err != nil {
		return errors.New("error when retrieving data " + err.Error())
	}
	resp.Animes = data
	return nil
}

func (s *AnimeService) GetAnimes(ctx context.Context, req *proto.Request, resp *proto.Results) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}
	query := map[string]interface{}{
		"title": req.GetTitle(),
		"type":  req.GetType(),
	}
	repo := s.GetRepo()
	defer repo.Close()

	data, err := repo.FindByQuery(query)
	if err != nil {
		return errors.New("error when retrieving data " + err.Error())
	}
	resp.Animes = data
	return nil
}
