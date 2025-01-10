package main

import (
	"log"
	"os"
	"reflect"

	config_application "github.com/RenzoReccio/golang-hexagonal-architecture/config/application"
	config_service "github.com/RenzoReccio/golang-hexagonal-architecture/config/service"
	"github.com/RenzoReccio/golang-hexagonal-architecture/presentation/controllers"
	base_controller "github.com/RenzoReccio/golang-hexagonal-architecture/presentation/controllers/base"

	"github.com/RenzoReccio/golang-hexagonal-architecture/presentation/middleware"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	controllers := discoverControllers()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	configService := config_service.NewConfigService()
	config_application.InitApplication(configService)
	registerControllers(router, controllers...)

	router.Run(":" + port)
}

func discoverControllers() []base_controller.BaseController {
	var controllersList []base_controller.BaseController

	pkgControllers := []interface{}{
		&controllers.WeatherController{},
	}

	for _, c := range pkgControllers {
		t := reflect.TypeOf(c).Elem()
		if t.Kind() == reflect.Struct {
			if reflect.PointerTo(t).Implements(reflect.TypeOf((*base_controller.BaseController)(nil)).Elem()) {
				controllersList = append(controllersList, reflect.New(t).Interface().(base_controller.BaseController))
			}
		}
	}

	return controllersList
}

func registerControllers(r *gin.Engine, controllers ...base_controller.BaseController) {
	for _, ctrl := range controllers {
		ctrl.RegisterRoutes(r)
	}
}
