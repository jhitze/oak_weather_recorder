package main

import (
	"fmt"
	"github.com/peterhellberg/sseclient"
	"log"
	"os"
	"strconv"
	"strings"
)

var logger *log.Logger
var urlFormat, deviceId, accessToken string

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	urlFormat = "https://api.particle.io/v1/devices/%s/events/?access_token=%s"

}

func main() {
	if len(os.Args) != 3 {
		logger.Println("Usage: OakWeatherRecorder device_id access_token")
		os.Exit(0)
	}
	deviceId = os.Args[1]
	accessToken = os.Args[2]

	url := fmt.Sprintf(urlFormat, deviceId, accessToken)
	events, err := sseclient.OpenURL(url)
	if err != nil {
		logger.Println("Error:", err)
		os.Exit(1)
	}
	logger.Printf("Connected to device %s", deviceId)
	for event := range events {
		// logger.Println(event.Name)
		data_decoded := NewWeatherData(event.Data)
		logger.Println(data_decoded.asString())
	}
}

type WeatherData struct {
	Temperature float64
	Ambient     float64
	Pressure    float64
	Humidity    float64
}

func (wd WeatherData) asString() string {

	return fmt.Sprintf("Temperature: %.2f°C, Ambient: %.2f%%, Pressure: %.2fmbar, Humidity: %.2f%%", wd.Temperature, wd.ambientPercentage(), wd.Pressure, wd.Humidity)
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
