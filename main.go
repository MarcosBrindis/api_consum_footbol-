package main

import (
	"consumer/src/consumer/infrastructure"
	"consumer/src/consumer/infrastructure/router"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar las dependencias necesarias para el consumo de mensajes
	infrastructure.InitDependencies()

	// Crear la instancia del router de Gin
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc:  func(origin string) bool { return true }, // Permite cualquier origen
		MaxAge:           12 * time.Hour,
	}))

	// Configurar las rutas para el consumidor
	router.SetupConsumerRoutes(r, infrastructure.ConsumeSportEventController)

	// Iniciar el servidor en el puerto 8080
	r.Run(":8001")
}
