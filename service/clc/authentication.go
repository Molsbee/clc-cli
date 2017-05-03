package clc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fmt"
	"github.com/molsbee/clc-cli/config"
	"github.com/molsbee/clc-cli/model"
)

// Authenticate performs a login request to CLC V2 Api with username and password provided
func (a *API) Authenticate(username string, password string) (err error) {
	authenticationRequest := &model.AuthenticationRequest{
		Username: username,
		Password: password,
	}
	authRequest, _ := json.Marshal(authenticationRequest)

	req, _ := http.NewRequest("POST", a.BaseURL+"/authentication/login", bytes.NewBuffer(authRequest))
	req.Header.Add("Content-Type", "application/json")

	resp, err := a.Client.Do(req)
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
