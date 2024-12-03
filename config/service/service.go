package config_service

import (
	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/repository"
	infraestructure_repository "github.com/RenzoReccio/golang-hexagonal-architecture/infrastructure/repository"
)

type ConfigService struct {
	WeatherRepository repository.IWeatherRepository
}

func NewConfigService() *ConfigService {
	connection := ""

	return &ConfigService{
		WeatherRepository: infraestructure_repository.NewWeatherRepository(&connection),
	}
}
