package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	updater "updater/proto/updater"
)

type Updater struct{}

func (e *Updater) Handle(ctx context.Context, msg *updater.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *updater.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
