package application_getweather

import (
	"context"
	"time"

	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/model"
	model_shared "github.com/RenzoReccio/golang-hexagonal-architecture/domain/model/shared"
	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/repository"
)

type GetWeatherHandler struct {
	weatherRepository repository.IWeatherRepository
}

func NewGetWorkItemTypeQueryHandler(weatherRepository repository.IWeatherRepository) *GetWeatherHandler {
	return &GetWeatherHandler{weatherRepository: weatherRepository}
}

func (c *GetWeatherHandler) Handle(ctx context.Context, query *GetWeatherQuery) (*model_shared.ResultWithValue[model.Weather], error) {

	result := c.weatherRepository.GetWeather()
	if !result.IsSuccess {
		return model_shared.NewResultWithValueFailure[model.Weather](model_shared.NewError("DB_ERROR", "Failure getting weather.")), nil
	}

	return model_shared.NewResultWithValueSuccess[model.Weather](model.NewWeather(time.Now(), 100)), nil
}
