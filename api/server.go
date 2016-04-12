package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/molsbee/clc-cli/authentication"
	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model"
)

// Server Exposes method for accessing Server Data
type Server struct {
	Name string
}

// Get Returns a Details object with Server Specific Details
func (s Server) Get() model.Details {
	req, _ := http.NewRequest("GET", constants.ClcAPIEndpoint+"/servers/"+authentication.Auth.AccountAlias+"/"+s.Name, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authentication.Auth.BearerToken)

	resp, _ := constants.Client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	res := model.Details{}
	json.Unmarshal(body, &res)

	if len(res.Details.IPAddresses) == 0 {
		log.Fatalln("No IP Addresses Returned for Server " + s.Name)
	}
	return res
}

// GetCredentials Returns Credentials object with Server Username and Password
func (s Server) GetCredentials() model.Credentials {
	req, _ := http.NewRequest("GET", constants.ClcAPIEndpoint+"/servers/"+authentication.Auth.AccountAlias+"/"+s.Name+"/credentials", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+authentication.Auth.BearerToken)

	resp, _ := constants.Client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	res := model.Credentials{}
	json.Unmarshal(body, &res)

	return res
}
