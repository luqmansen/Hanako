package handler

import (
	"context"
	"errors"
	"github.com/luqmansen/hanako/services/anime/handler/datastore"
	proto "github.com/luqmansen/hanako/services/anime/proto"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/trace"
	"gopkg.in/mgo.v2"
)

//AnimeService is type for containing the session
type AnimeService struct {
	Session *mgo.Session
}

//GetRepo return a session from anime repository connection
func (s *AnimeService) GetRepo() datastore.Repository {
	return &datastore.AnimeRepository{Session: s.Session.Clone()}
}

//GetAll is implementation of rpc AnimeService Interface
func (s *AnimeService) GetAll(ctx context.Context, req *proto.Request, resp *proto.Results) error {
	md, _ := metadata.FromContext(ctx)
	traceID := md["traceID"]
	if tr, ok := trace.FromContext(ctx); ok {
		tr.LazyPrintf("traceID %s", traceID)
	}
	wireContext, _ := opentracing.GlobalTracer().Extract(opentracing.TextMap, opentracing.TextMapCarrier(md))
	sp := opentracing.StartSpan("anime-srv", opentracing.ChildOf(wireContext))
	sp.SetTag("req", req)
	defer func() {
		sp.SetTag("res", resp)
		sp.Finish()
	}()

	repo := s.GetRepo()
	defer repo.Close()

	data, err := repo.FindAll()
	if err != nil {
		return errors.New("error when retrieving data " + err.Error())
	}
	resp.Animes = data
	return nil
}

//GetAnimes is also implementation of rpc AnimeService Interface
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
