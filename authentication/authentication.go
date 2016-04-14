package authentication

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/molsbee/clc-cli/config"
	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model"
)

// Auth AuthenticationResponse Store on Local FileSystem
var Auth = model.AuthenticationResponse{}

func init() {
	fileData, _ := ioutil.ReadFile(config.FilePath())
	json.Unmarshal(fileData, &Auth)
}

// Authenticate Authenticate with CLC API
func Authenticate(username string, password string) {
	authenticationRequest := &model.AuthenticationRequest{
		Username: username,
		Password: password,
	}
	authRequest, _ := json.Marshal(authenticationRequest)

	req, _ := http.NewRequest("POST", constants.ClcAPIEndpoint+"/authentication/login", bytes.NewBuffer(authRequest))
	req.Header.Add("Content-Type", "application/json")

	resp, _ := constants.Client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Fatalln("Failed to Authenticate with API")
	}

	ioutil.WriteFile(config.FilePath(), body, 0777)
}
