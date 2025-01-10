package controllers

import (
	"net/http"

	application_createweather "github.com/RenzoReccio/golang-hexagonal-architecture/application/weather/command/create-weather"
	"github.com/RenzoReccio/golang-hexagonal-architecture/domain/model"
	model_shared "github.com/RenzoReccio/golang-hexagonal-architecture/domain/model/shared"
	"github.com/gin-gonic/gin"
	"github.com/mehdihadeli/go-mediatr"
)

type WeatherController struct {
}

func (u *WeatherController) RegisterRoutes(r *gin.Engine) {
	userGroup := r.Group("/weather")
	{
		userGroup.POST("/", u.InsertWeather)
	}
}

func (u *WeatherController) InsertWeather(c *gin.Context) {
	req := new(application_createweather.CreateWeatherCommand)

	if err := c.Bind(req); err != nil {
		c.IndentedJSON(http.StatusBadRequest, model_shared.NewResultFailure(model_shared.NewError("BAD_REQUEST", "Input not correct format")))
		return
	}
	result, _ := mediatr.Send[*application_createweather.CreateWeatherCommand, *model_shared.ResultWithValue[model.Weather]](c, req)

	c.IndentedJSON(http.StatusCreated, result)
}
