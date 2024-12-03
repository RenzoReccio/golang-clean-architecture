package application_sendweather

import (
	"context"
	"errors"
	"time"

	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/model"
	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/repository"
)

type SendWeatherHandler struct {
	weatherRepository repository.IWeatherRepository
}

func NewSendWeatherHandler(weatherRepository repository.IWeatherRepository) *SendWeatherHandler {
	return &SendWeatherHandler{weatherRepository: weatherRepository}
}

func (c *SendWeatherHandler) Handle(ctx context.Context, weather *SendWeatherEvent) error {
	result := c.weatherRepository.SendWeather(*model.NewWeather(time.Now(), weather.Temperature))
	if !result.IsSuccess {
		return errors.New(result.Error.Description)
	}
	return nil
}
