package infrastructure

import (
	"consumer/src/consumer/application"
	"consumer/src/consumer/infrastructure/http/controller"
	"consumer/src/consumer/infrastructure/messaging"
	"fmt"
)

var (
	ConsumeSportEventController *controller.ConsumeSportEventController
)

func InitDependencies() {
	// Inicializar conexión a RabbitMQ para consumir mensajes
	rabbitMQAdapter, err := messaging.InitRabbitMQ()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to RabbitMQ: %v", err))
	}

	consumeUsecase := application.NewConsumeSportsEventUsecase(rabbitMQAdapter)
	ConsumeSportEventController = controller.NewConsumeSportEventController(consumeUsecase)
}
