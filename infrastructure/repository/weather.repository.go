package infraestructure_repository

import (
	model "github.com/RenzoReccio/golang-hexagonal-architecture/domain/model"
	model_shared "github.com/RenzoReccio/golang-hexagonal-architecture/domain/model/shared"
	repository "github.com/RenzoReccio/golang-hexagonal-architecture/domain/repository"
)

type WeatherRepository struct {
	connection *string
}

func NewWeatherRepository(connection *string) repository.IWeatherRepository {
	return &WeatherRepository{
		connection: connection,
	}
}

func (u *WeatherRepository) GetWeather() *model_shared.ResultWithValue[[]model.Weather] {
	weathers := []model.Weather{}
	return model_shared.NewResultWithValueSuccess[[]model.Weather](&weathers)
}
func (u *WeatherRepository) InsertWeather(weather model.Weather) *model_shared.Result {
	return model_shared.NewResultSuccess()
}

func (u *WeatherRepository) SendWeather(weather model.Weather) *model_shared.Result {
	return model_shared.NewResultSuccess()
}
