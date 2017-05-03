package clc

import (
	"github.com/molsbee/clc-cli/model"
)

type GroupRequest struct {
	AccountAlias string
	ID           string
}

// GetGroup returns a Group struct based on the data provided.  If AccountAlias isn't set then the authentication details
// will be used for the account alias.
func (a *API) GetGroup(request GroupRequest) model.Group {
	endpoint := a.BaseURL + "/groups/" + a.GetOrDefaultAccountAlias(request.AccountAlias) + "/" + request.ID

	group := model.Group{}
	a.GetWithAuthentication(endpoint, &group)

	return group
}
