package server

import (
	"testing"
)

func TestDeployOk(t *testing.T) {
	weather := WeatherCondition{
		Temperature: 20,
		WindSpeed:   10,
	}
	deploy := CheckWeather(weather)
	if !deploy {
		t.Errorf("Got deploy=%v. Excepted true", deploy)
	}
}

func TestTempError(t *testing.T) {
	weather := WeatherCondition{
		Temperature: 30,
		WindSpeed:   20,
	}
	deploy := CheckWeather(weather)
	if deploy {
		t.Errorf("Got deploy=%v. Excepted false", deploy)
	}

	weather.Temperature = -15
	deploy = CheckWeather(weather)
	if deploy {
		t.Errorf("Got deploy=%v. Excepted false", deploy)
	}
}
func TestWindError(t *testing.T) {
	weather := WeatherCondition{
		Temperature: 20,
		WindSpeed:   40,
	}
	deploy := CheckWeather(weather)
	if deploy {
		t.Errorf("Got deploy=%v. Excepted false", deploy)
	}
}
