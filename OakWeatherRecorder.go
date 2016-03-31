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
	var username, password string
	fmt.Print("Enter username: ")
	fmt.Scanln(&username)
	fmt.Print("Enter password: ")
	fmt.Scanln(&password)

	accessToken := retrieveAccessToken(username, password)
	deviceList := retrieveDevices(accessToken)
	device := askWhichDevice(deviceList)
	listenForWeatherEvents(device, accessToken)
}

func listenForWeatherEvents(device Device, accessToken string) {
	url := fmt.Sprintf(urlFormat, device.Id, accessToken)
	events, err := sseclient.OpenURL(url)
	if err != nil {
		logger.Println("Error:", err)
		os.Exit(1)
	}
	logger.Printf("Connected to the stream of device %s (%s)", device.Name, device.Id)
	for event := range events {
		data_decoded := NewWeatherData(event.Data)
		logger.Println(data_decoded.asString())
	}
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
