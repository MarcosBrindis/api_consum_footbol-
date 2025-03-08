// application/consume_sports_event_usecase.go
package application

import (
	"consumer/src/consumer/domain/ports"
	"context"
)

type ConsumeSportsEventUsecase struct {
	Consumer ports.EventConsumer
}

func NewConsumeSportsEventUsecase(consumer ports.EventConsumer) *ConsumeSportsEventUsecase {
	return &ConsumeSportsEventUsecase{
		Consumer: consumer,
	}
}

func (uc *ConsumeSportsEventUsecase) Execute(ctx context.Context) (string, error) {
	return uc.Consumer.GetMessage(ctx)
}
