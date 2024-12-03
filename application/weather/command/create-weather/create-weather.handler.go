package application_createweather

import (
	"context"
	"time"

	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/model"
	model_shared "github.com/RenzoReccio/golang-hexagonal-architecture/domain/model/shared"
	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/repository"
)

type CreateWeatherHandler struct {
	weatherRepository repository.IWeatherRepository
}

func NewCreateWeatherCommandHandler(weatherRepository repository.IWeatherRepository) *CreateWeatherHandler {
	return &CreateWeatherHandler{weatherRepository: weatherRepository}
}

func (c *CreateWeatherHandler) Handle(ctx context.Context, command *CreateWeatherCommand) (*model_shared.ResultWithValue[model.Weather], error) {

	result := c.weatherRepository.InsertWeather(*model.NewWeather(time.Now(), command.Temperature))
	if !result.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Weather](model_shared.NewError("DB_ERROR", "Failure inserting weather.")), nil
	}

	return model_shared.NewResultWithValueSuccess[model.Weather](model.NewWeather(time.Now(), 100)), nil
}
