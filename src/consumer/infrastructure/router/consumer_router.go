package router

import (
	"consumer/src/consumer/infrastructure/http/controller"

	"github.com/gin-gonic/gin"
)

func SetupConsumerRoutes(
	r *gin.Engine,
	consume *controller.ConsumeSportEventController) {
	consumerGroup := r.Group("/consumer")
	{
		consumerGroup.GET("/message", consume.HandleConsume)
	}
}
