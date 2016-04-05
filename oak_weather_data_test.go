package main

import (
	"testing"
)

func TestWeatherData_asString(t *testing.T) {
	weatherData := WeatherData{24.5, 128, 1048.23, 56}
	expected := "Temperature: 24.50°C, Ambient: 50.00%, Pressure: 1048.23 mbar, Humidity: 56.00%"
	returned := weatherData.asString()

	if expected != returned {
		t.Errorf("WeatherData#asString did not format correctly. Got:%s", returned)
	}
}

func TestWeatherData_ambientPercentage_Middle(t *testing.T) {
	weatherData := WeatherData{24.5, 128, 1048.23, 56}
	expected := 50.0
	returned := weatherData.ambientPercentage()

	if expected != returned {
		t.Errorf("WeatherData#ambientPercentage did return correct percentage. Got:%f", returned)
	}
}

func TestWeatherData_ambientPercentage_High(t *testing.T) {
	weatherData := WeatherData{24.5, 10, 1048.23, 56}
	expected := 96.09375
	returned := weatherData.ambientPercentage()

	if expected != returned {
		t.Errorf("WeatherData#ambientPercentage did return correct percentage. Got:%f", returned)
	}
}

func TestWeatherData_ambientPercentage_Low(t *testing.T) {
	weatherData := WeatherData{24.5, 200, 1048.23, 56}
	expected := 21.875
	returned := weatherData.ambientPercentage()

	if expected != returned {
		t.Errorf("WeatherData#ambientPercentage did return correct percentage. Got:%f", returned)
	}
}
