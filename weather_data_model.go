package main

import (
	"fmt"
	"strconv"
	"strings"
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

func NewWeatherData(data map[string]interface{}) WeatherData {
	var weatherData WeatherData
	for key, value := range data {
		if key == "data" {
			weather_variables := strings.Split(value.(string), ";")
			for _, variable := range weather_variables {
				split_weather_variable := strings.Split(variable, "-")
				switch split_weather_variable[0] {
				case "temp":
					weatherData.Temperature, _ = strconv.ParseFloat(split_weather_variable[1], 64)
				case "ambient":
					weatherData.Ambient, _ = strconv.ParseFloat(split_weather_variable[1], 64)
				case "pressue":
					weatherData.Pressure, _ = strconv.ParseFloat(split_weather_variable[1], 64)
				case "humidity":
					weatherData.Humidity, _ = strconv.ParseFloat(split_weather_variable[1], 64)
				case "weather":
					break
				default:
					logger.Println("Ignored:", split_weather_variable)
				}

			}
		}
	}
	return weatherData
}
