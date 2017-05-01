package api

import (
	"log"

	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model"
	"github.com/molsbee/clc-cli/net"
)

// Server Exposes method for accessing Server Data
type Server struct {
	BaseRequest
	Name string
}

// Get Returns a Details object with Server Specific Details
func (s Server) Get() model.Details {
	endpoint := constants.ClcAPIEndpoint + "/servers/" + s.GetAccountAlias() + "/" + s.Name

	details := model.Details{}
	net.GetWithAuthentication(endpoint, &details)

	if len(details.Details.IPAddresses) == 0 {
		log.Fatalln("No IP Addresses Returned for Server " + s.Name)
	}
	return details
}

// GetCredentials Returns Credentials object with Server Username and Password
func (s Server) GetCredentials() model.Credentials {
	endpoint := constants.ClcAPIEndpoint + "/servers/" + s.GetAccountAlias() + "/" + s.Name + "/credentials"

	credentials := model.Credentials{}
	net.GetWithAuthentication(endpoint, &credentials)

	return credentials
}
