package clc

import (
	"github.com/molsbee/clc-cli/model"
)

// Server Exposes method for accessing Server Data
type ServerRequest struct {
	AccountAlias string
	Name         string
}

// GetServer Returns a Details object with Server Specific Details
func (a *API) GetServer(request ServerRequest) model.Details {
	endpoint := a.BaseURL + "/servers/" + a.GetOrDefaultAccountAlias(request.AccountAlias) + "/" + request.Name

	details := model.Details{}
	a.GetWithAuthentication(endpoint, &details)

	return details
}

// GetServerCredentials Returns Credentials object with Server Username and Password
func (a *API) GetServerCredentials(request ServerRequest) model.Credentials {
	endpoint := a.BaseURL + "/servers/" + a.GetOrDefaultAccountAlias(request.AccountAlias) + "/" + request.Name + "/credentials"

	credentials := model.Credentials{}
	a.GetWithAuthentication(endpoint, &credentials)

	return credentials
}
