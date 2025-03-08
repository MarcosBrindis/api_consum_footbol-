package main

import (
	"consumer/src/consumer/infrastructure"
	"consumer/src/consumer/infrastructure/router"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar las dependencias necesarias para el consumo de mensajes
	infrastructure.InitDependencies()

	// Crear la instancia del router de Gin
	r := gin.Default()

	// Configurar las rutas para el consumidor
	router.SetupConsumerRoutes(r, infrastructure.ConsumeSportEventController)

	// Iniciar el servidor en el puerto 8080
	r.Run(":8001")
}
