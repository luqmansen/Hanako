package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	search "search/proto/search"
)

type Search struct{}

func (e *Search) Handle(ctx context.Context, msg *search.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *search.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
