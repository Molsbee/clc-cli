package server

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/molsbee/clc-cli/authentication"
	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model"
)

// GetDetails Returns a Details object with Server Specific Details
func GetDetails(server string) model.Details {
	req, _ := http.NewRequest("GET", constants.ClcAPIEndpoint+"/servers/"+authentication.Auth.AccountAlias+"/"+server, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authentication.Auth.BearerToken)

	resp, _ := constants.Client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	res := model.Details{}
	json.Unmarshal(body, &res)

	if len(res.Details.IPAddresses) == 0 {
		log.Fatalln("No IP Addresses Returned for Server " + server)
	}
	return res
}

// GetCredentials Returns Credentials object with Server Username and Password
func GetCredentials(server string) model.Credentials {
	req, _ := http.NewRequest("GET", constants.ClcAPIEndpoint+"/servers/"+authentication.Auth.AccountAlias+"/"+server+"/credentials", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authentication.Auth.BearerToken)

	resp, _ := constants.Client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	res := model.Credentials{}
	json.Unmarshal(body, &res)

	return res
}
