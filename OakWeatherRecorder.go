package main

import (
	"fmt"
	"github.com/peterhellberg/sseclient"
	"log"
	"os"
)

var logger *log.Logger
var urlFormat, deviceId string

func init() {
	logger = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	urlFormat = "https://api.particle.io/v1/devices/%s/events/?access_token=%s"

}

func main() {
	settings := &OakWeatherSettings{}
	var err error
	settings, err = findSettings()
	if err != nil {
		logger.Println("Could not read settings file. reason:", err)
		logger.Println("Reverting to asking for the settings.")
		settings, err = askForSettings()
		saveSettings(*settings)
	}

	events_channel := openConnectionForWeatherEvents(settings.SelectedDevice, settings.AccessToken)
	listenForWeatherEvents(events_channel)
}

func openConnectionForWeatherEvents(device Device, accessToken string) chan sseclient.Event {
	url := fmt.Sprintf(urlFormat, device.Id, accessToken)
	events, err := sseclient.OpenURL(url)
	if err != nil {
		logger.Println("Error:", err)
		os.Exit(1)
	}
	logger.Printf("Connected to the stream of device %s (%s)", device.Name, device.Id)
	return events
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
