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

func TestNewWeatherData_withGoodData(t *testing.T) {
	weatherJsonMap := make(map[string]interface{})
	weatherData := "{\"temperature\" : 24.5, \"ambient\" : 200, \"pressure\" : 1048.23, \"humidity\" : 56}"

	weatherJsonMap["data"] = weatherData

	expected := WeatherData{24.5, 200, 1048.23, 56}

	returned := NewWeatherData(weatherJsonMap)

	if expected != *returned {
		t.Errorf("TestNewWeatherData did return correct data. Got:%s, expected %s", returned.asString(), expected.asString())
	}
}

func TestNewWeatherData_withNoData(t *testing.T) {
	weatherJsonMap := make(map[string]interface{})

	returned := NewWeatherData(weatherJsonMap)

	if returned != nil {
		t.Error("TestNewWeatherData did not return nil pointer.")
	}
}

func TestNewWeatherData_withBadJSONinData(t *testing.T) {
	weatherJsonMap := make(map[string]interface{})
	weatherData := "\"temperature\" : 24.5, \"ambient\" : 200, \"pressure\" : 1048.23, \"humidity\" : 56"

	weatherJsonMap["data"] = weatherData

	returned := NewWeatherData(weatherJsonMap)

	if returned != nil {
		t.Error("TestNewWeatherData did not return nil pointer.")
	}
}

func TestNewWeatherData_withBlankData(t *testing.T) {
	weatherJsonMap := make(map[string]interface{})
	weatherJsonMap["data"] = ""
	returned := NewWeatherData(weatherJsonMap)

	if returned != nil {
		t.Error("TestNewWeatherData did not return nil pointer.")
	}
}
