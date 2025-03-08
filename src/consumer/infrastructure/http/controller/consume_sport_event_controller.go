package controller

import (
	"context"
	"net/http"

	"consumer/src/consumer/application"

	"github.com/gin-gonic/gin"
)

type ConsumeSportEventController struct {
	UseCase *application.ConsumeSportsEventUsecase
}

func NewConsumeSportEventController(useCase *application.ConsumeSportsEventUsecase) *ConsumeSportEventController {
	return &ConsumeSportEventController{
		UseCase: useCase,
	}
}

func (c *ConsumeSportEventController) HandleConsume(ctx *gin.Context) {
	// Puedes usar ctx.Request.Context() para propagar el contexto si lo prefieres
	reqCtx := context.Background()
	message, err := c.UseCase.Execute(reqCtx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": message})
}
