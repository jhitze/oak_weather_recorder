package main

import (
	"encoding/json"
	"fmt"
)

type WeatherData struct {
	Temperature float64
	Ambient     float64
	Pressure    float64
	Humidity    float64
}

func (wd WeatherData) asString() string {

	return fmt.Sprintf("Temperature: %.2f°C, Ambient: %.2f%%, Pressure: %.2f mbar, Humidity: %.2f%%", wd.Temperature, wd.ambientPercentage(), wd.Pressure, wd.Humidity)
}

func (wd WeatherData) ambientPercentage() float64 {
	return (256 - wd.Ambient) / 256 * 100
}

func NewWeatherData(data map[string]interface{}) *WeatherData {
	var weatherData WeatherData
	for key, value := range data {
		if key == "data" {
			jsonerr := json.Unmarshal([]byte(value.(string)), &weatherData)
			if jsonerr != nil {
				logger.Println("Error unmarshalling the weatherJSON response:", jsonerr)
				logger.Println("Server sent:", value)
				return nil
			}
		}
	}
	return &weatherData
}
