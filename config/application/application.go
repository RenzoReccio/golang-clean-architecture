package config_application

import (
	application_createweather "github.com/RenzoReccio/golang-hexagonal-architecture/application/weather/command/create-weather"
	application_sendweather "github.com/RenzoReccio/golang-hexagonal-architecture/application/weather/event/send-weather"
	application_getweather "github.com/RenzoReccio/golang-hexagonal-architecture/application/weather/query/get-weather"
	config_service "github.com/RenzoReccio/golang-hexagonal-architecture/config/service"
	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/model"
	model_shared "github.com/RenzoReccio/golang-hexagonal-architecture/domain/model/shared"
	"github.com/mehdihadeli/go-mediatr"
)

func InitApplication(configService *config_service.ConfigService) {

	getWeatherQueryHandler := application_getweather.NewGetWeatherHandler(configService.WeatherRepository)
	createWeatherQueryHandler := application_createweather.NewCreateWeatherCommandHandler(configService.WeatherRepository)
	sendWeatherEvent := application_sendweather.NewSendWeatherHandler(configService.WeatherRepository)

	mediatr.RegisterRequestHandler[*application_createweather.CreateWeatherCommand, *model_shared.ResultWithValue[model.Weather]](createWeatherQueryHandler)
	mediatr.RegisterRequestHandler[*application_getweather.GetWeatherQuery, *model_shared.ResultWithValue[model.Weather]](getWeatherQueryHandler)
	mediatr.RegisterNotificationHandlers[*application_sendweather.SendWeatherEvent](sendWeatherEvent)
}
