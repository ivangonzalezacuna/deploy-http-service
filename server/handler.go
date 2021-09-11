package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type DeployStatus struct {
	Deploy      bool   `json:"deploy"`
	Temperature int    `json:"current_temp"`
	Wind        int    `json:"current_wind"`
	Error       string `json:"error"`
	CuttieFox   string `json:"cutie_fox"`
}

type WeatherInfo struct {
	CurrentCondition []WeatherCondition `json:"current_condition"`
}

type WeatherCondition struct {
	Temperature int `json:"temp_C,string"`
	WindSpeed   int `json:"windspeedKmph,string"`
}

type FoxData struct {
	Image string `json:"image"`
}

// GetWeatherCondition makes a request to get all the weather
// information in Leipzig and unmarshals the necesssary information
func GetWeatherCondition() (*WeatherCondition, error) {
	res, err := http.Get("https://wttr.in/Leipzig?format=j1")
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	weatherCondition := WeatherInfo{}
	json.Unmarshal(responseData, &weatherCondition)
	currentWeather := weatherCondition.CurrentCondition[0]
	return &currentWeather, nil
}

// GetFoxLink makes a request to an API to get a link
// to a fox image
func GetFoxLink() string {
	res, err := http.Get("https://randomfox.ca/floof/")
	if err != nil {
		return ""
	}

	responseData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ""
	}

	fox := FoxData{}
	json.Unmarshal(responseData, &fox)

	return fox.Image
}

// CheckWeather receives the current weather condition and
// evaluates if all the conditions to deploy the code are
// satisfied or not
func CheckWeather(weather WeatherCondition) bool {
	fmt.Printf("Current weather: %+v\n", weather)
	tempOk, windOk := true, true

	if weather.Temperature < -10 || weather.Temperature > 25 {
		tempOk = false
	}
	if weather.WindSpeed > 20 {
		windOk = false
	}

	return tempOk && windOk
}

// CheckDeploy makes a request for the current weather condition
// and determines if the temperature and wind speed is good enough
// to deploy all your code
func CheckDeploy() DeployStatus {
	var deploy DeployStatus

	currentWeather, err := GetWeatherCondition()
	if err != nil {
		deploy.Error = "Error getting weather condition"
		return deploy
	}

	status := CheckWeather(*currentWeather)

	deploy.Temperature = currentWeather.Temperature
	deploy.Wind = currentWeather.WindSpeed
	deploy.Deploy = status
	deploy.CuttieFox = GetFoxLink()

	return deploy
}
