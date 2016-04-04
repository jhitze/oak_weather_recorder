package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/peterhellberg/sseclient"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
)

func retrieveAccessToken(username, password string) string {
	var oauth_url = "https://api.particle.io/oauth/token"

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)

	req, err := http.NewRequest("POST",
		oauth_url,
		bytes.NewBufferString(data.Encode()))

	// this is only for use by single developer, not an application.
	req.SetBasicAuth("particle", "particle")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Printf("Got an error -> %s", err)
	}

	var oathResponse OauthTokenResponse
	jsonerr := json.Unmarshal(body, &oathResponse)
	if jsonerr != nil {
		logger.Println("Error unmarshalling the oath response:", jsonerr)
		return ""
	}

	return oathResponse.Access_token
}

func retrieveDevices(accessToken string) []Device {
	var devices_url = "https://api.particle.io/v1/devices"

	req, err := http.NewRequest("GET",
		devices_url,
		nil)

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Printf("Got an error -> %s", err)
	}

	var deviceList []Device
	jsonerr := json.Unmarshal(body, &deviceList)
	if jsonerr != nil {
		logger.Println("Error unmarshalling the device list:", jsonerr)
		return nil
	}

	return deviceList
}

func retrieveDevice(accessToken, deviceId string) (*Device, error) {
	logger.Println("Getting current information for device:", deviceId)
	device := Device{}
	var devices_url = "https://api.particle.io/v1/devices/" + deviceId

	req, err := http.NewRequest("GET",
		devices_url,
		nil)

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	jsonerr := json.Unmarshal(body, &device)
	if jsonerr != nil {
		return nil, jsonerr
	}

	return &device, nil
}

func openConnectionForWeatherEvents(device Device, accessToken string) chan sseclient.Event {
	urlFormat = "https://api.particle.io/v1/devices/%s/events/?access_token=%s"
	url := fmt.Sprintf(urlFormat, device.Id, accessToken)
	events, err := sseclient.OpenURL(url)
	if err != nil {
		logger.Println("Error:", err)
		os.Exit(1)
	}
	logger.Printf("Connected to the stream of device %s (%s)", device.Name, device.Id)
	return events
}

type OauthTokenResponse struct {
	Access_token  string
	Token_type    string
	Expires_in    int
	Refresh_token string
}

type Device struct {
	Id              string
	Name            string
	Last_app        string
	Last_ip_address string
	Last_heard      string
	Product_id      float64
	Connected       bool
}
