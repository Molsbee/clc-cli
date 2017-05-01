package api

import (
	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model"
	"github.com/molsbee/clc-cli/net"
)

// Group struct represent the necessary data to perform Group API request to CLC V2 API.
type Group struct {
	BaseRequest
	ID string
}

// Get returns a Group struct based on the data provided.  If AccountAlias isn't set then the authentication details
// will be used for the account alias.
func (g Group) Get() model.Group {
	endpoint := constants.ClcAPIEndpoint + "/groups/" + g.GetAccountAlias() + "/" + g.ID

	group := model.Group{}
	net.GetWithAuthentication(endpoint, &group)

	return group
}
