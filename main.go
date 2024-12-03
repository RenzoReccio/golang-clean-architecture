package main

import (
	"log"
	"os"

	config_application "github.com/RenzoReccio/golang-hexagonal-architecture/config/application"
	config_service "github.com/RenzoReccio/golang-hexagonal-architecture/config/service"
	"github.com/RenzoReccio/golang-hexagonal-architecture/presentation/controllers"
	"github.com/RenzoReccio/golang-hexagonal-architecture/presentation/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initializeReceiverHttpHandler(router *gin.Engine) {
	// Initialize all layers
	weatherController := controllers.NewWeatherController()
	// Routers
	router.POST("weather", weatherController.InsertWeather)
}

func main() {
	godotenv.Load(".env")
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	configService := config_service.NewConfigService()
	config_application.InitApplication(configService)
	initializeReceiverHttpHandler(router)
	router.Run(":" + port)
}
