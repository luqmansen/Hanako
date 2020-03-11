package handler

import (
	"context"
	"github.com/luqmansen/hanako/api/rest/responses"
	r "github.com/luqmansen/hanako/api/rest/utils"
	"github.com/luqmansen/hanako/services/anime/proto"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/micro/api/proto"
)

type Anime struct {
	Client client.Client
}

type animeResults struct {
	animeList []*anime.Anime
	err       error
}
type querySearch struct {
	title string
	kind  string
}

func (a *Anime) Search(ctx context.Context, req *go_micro_api.Request, rsp *go_micro_api.Response) error {
	const function = ".search"
	q := &querySearch{
		title: getUrlParam("title", req.Get),
		kind:  getUrlParam("type", req.Get),
	}

	animeCh := queryAnime(a.Client, ctx, q)
	animeReply := <-animeCh
	if err := animeReply.err; err != nil {
		return errors.InternalServerError(r.ServiceName+function, err.Error())
	}
	b, s := responses.JSON([]interface{}{animeReply.animeList})
	rsp.StatusCode = s
	rsp.Body = b

	return nil
}

func queryAnime(c client.Client, ctx context.Context, q *querySearch) chan animeResults {
	animeClient := anime.NewAnimeService("hanako.srv.anime", c)
	ch := make(chan animeResults, 1)

	go func() {
		res, err := animeClient.GetAnimes(ctx, &anime.Request{Title: q.title, Type: q.kind})
		if res != nil {
			ch <- animeResults{res.Animes, err}
		}
	}()
	return ch
}

func getUrlParam(key string, p map[string]*go_micro_api.Pair) string {
	var v string
	if val, ok := p[key]; ok {
		v = val.Values[0]
	}
	return v
}
