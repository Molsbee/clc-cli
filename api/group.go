package api

import (
	"github.com/molsbee/clc-cli/authentication"
	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model"
	"github.com/molsbee/clc-cli/net"
)

// Group Stuff
type Group struct {
	ID string
}

// Get Stuff
func (g Group) Get() model.Group {
	endpoint := constants.ClcAPIEndpoint + "/groups/" + authentication.Auth.AccountAlias + "/" + g.ID

	group := model.Group{}
	net.GetWithAuthentication(endpoint, &group)

	return group
}
