package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func askForSettings() (*OakWeatherSettings, error) {
	settings := OakWeatherSettings{}
	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	settings.AccessToken = retrieveAccessToken(username, password)
	deviceList := retrieveDevices(settings.AccessToken)
	settings.SelectedDevice = askWhichDevice(deviceList)
	return &settings, nil
}

func findSettings() (*OakWeatherSettings, error) {
	settings := OakWeatherSettings{}
	filename := "oak_weather.json"
	logger.Println("Going to attempt to load data from", filename)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &settings)
	if err != nil {
		return nil, err
	}

	device := &Device{}
	device, err = retrieveDevice(settings.AccessToken, settings.SelectedDevice.Id)
	settings.SelectedDevice = *device
	return &settings, nil
}

func askWhichDevice(deviceList []Device) Device {
	var deviceNumber int
	fmt.Println("Found Devices")
	fmt.Println("----------------------------------------")
	for number, device := range deviceList {
		fmt.Printf("%d: %s - Online->%t\n", number, device.Name, device.Connected)
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Pick a number:")
	fmt.Scanln(&deviceNumber)
	logger.Printf("Device %s picked.", deviceList[deviceNumber].Name)
	return deviceList[deviceNumber]
}

type OakWeatherSettings struct {
	SelectedDevice Device
	AccessToken    string
}