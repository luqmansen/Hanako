package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	updater "updater/proto/updater"
)

type Updater struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Updater) Call(ctx context.Context, req *updater.Request, rsp *updater.Response) error {
	log.Log("Received Updater.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Updater) Stream(ctx context.Context, req *updater.StreamingRequest, stream updater.Updater_StreamStream) error {
	log.Logf("Received Updater.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&updater.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Updater) PingPong(ctx context.Context, stream updater.Updater_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&updater.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
