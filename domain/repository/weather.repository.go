package repository

import (
	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/model"
	model_shared "github.com/RenzoReccio/golang-hexagonal-architecture/domain/model/shared"
)

type IWeatherRepository interface {
	GetWeather() *model_shared.ResultWithValue[[]model.Weather]
	InsertWeather(weather model.Weather) *model_shared.Result
	SendWeather(weather model.Weather) *model_shared.Result
}
