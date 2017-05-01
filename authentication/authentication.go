package authentication

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/molsbee/clc-cli/config"
	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model"
	"fmt"
)

// Auth AuthenticationResponse Store on Local FileSystem
var Auth = model.AuthenticationResponse{}

func init() {
	fileData, _ := ioutil.ReadFile(config.FilePath())
	json.Unmarshal(fileData, &Auth)
}

// Authenticate performs a login request to CLC V2 Api with username and password provided
func Authenticate(username string, password string) (err error) {
	authenticationRequest := &model.AuthenticationRequest{
		Username: username,
		Password: password,
	}
	authRequest, _ := json.Marshal(authenticationRequest)

	req, _ := http.NewRequest("POST", constants.ClcAPIEndpoint+"/authentication/login", bytes.NewBuffer(authRequest))
	req.Header.Add("Content-Type", "application/json")

	resp, err := constants.Client.Do(req)
	if err != nil {
		err = fmt.Errorf("Authentiication request failed - %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil || resp.StatusCode != 200 {
		err = fmt.Errorf("Failed to authenticate user %s", username)
	}

	if err == nil {
		ioutil.WriteFile(config.FilePath(), body, 0777)
	}
	return
}
