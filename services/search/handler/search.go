package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	search "search/proto/search"
)

type Search struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Search) Call(ctx context.Context, req *search.Request, rsp *search.Response) error {
	log.Log("Received Search.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Search) Stream(ctx context.Context, req *search.StreamingRequest, stream search.Search_StreamStream) error {
	log.Logf("Received Search.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&search.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Search) PingPong(ctx context.Context, stream search.Search_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&search.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
