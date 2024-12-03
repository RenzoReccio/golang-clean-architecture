package model

import "time"

type Weather struct {
	Date        time.Time
	Temperature int
}

func NewWeather(date time.Time, temperature int) *Weather {
	return &Weather{
		Date:        date,
		Temperature: temperature,
	}
}
