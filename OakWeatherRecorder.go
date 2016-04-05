package main

import (
	"github.com/peterhellberg/sseclient"
	"log"
	"os"
)

var logger *log.Logger
var urlFormat, deviceId string

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)

}

func main() {
	settings := loadSettings()
	events_channel := openConnectionForWeatherEvents(settings.SelectedDevice, settings.AccessToken)
	listenForWeatherEvents(events_channel)
}

func listenForWeatherEvents(events chan sseclient.Event) {
	for event := range events {
		if event.Name == "weatherstationJSON" {
			data_decoded := NewWeatherData(event.Data)
			if data_decoded != nil {
				logger.Println(data_decoded.asString())
			}
		}
	}
}
