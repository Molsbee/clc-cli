package net

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/molsbee/clc-cli/authentication"
)

// Client Global Http Client
var client = &http.Client{}

// GetWithAuthentication Perform CLC API Get Request with BearerToken Authentication
func GetWithAuthentication(endpoint string, resp interface{}) {
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authentication.Auth.BearerToken)

	response, err := client.Do(req)
	if err != nil {
		log.Fatalln("Errer received while performing clc api request", err)
	}
	defer response.Body.Close()
	if response.StatusCode < 200 || response.StatusCode > 299 {
		log.Fatalf("Non %d status code received from api\n", response.StatusCode)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln("Unable to read response body")
	}

	json.Unmarshal(body, resp)
}
