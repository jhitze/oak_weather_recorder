package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func retrieveAccessToken(username, password string) string {
	var oauth_url = "https://api.particle.io/oauth/token"

	data := url.Values{}
	data.Set("grant_type", "password")
	data.Set("username", username)
	data.Set("password", password)
	data.Set("expires_in", "10")

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
