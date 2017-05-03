package clc

import (
	"github.com/molsbee/clc-cli/model"
)

type DataCenterRequest struct {
	AccountAlias string
	Name         string
}

// GetDataCenter Returns Data Center details
func (a *API) GetDataCenter(request DataCenterRequest) model.DataCenter {
	accountAlias := a.GetOrDefaultAccountAlias(request.AccountAlias)
	endpoint := a.BaseURL + "/datacenters/" + accountAlias + "/" + request.Name + "?groupLinks=true"

	dataCenter := model.DataCenter{}
	a.GetWithAuthentication(endpoint, &dataCenter)

	return dataCenter
}

// GetDataCenters Returns List of Data Centers
func (a *API) GetDataCenters(request DataCenterRequest) model.DataCenterList {
	accountAlias := a.GetOrDefaultAccountAlias(request.AccountAlias)
	endpoint := a.BaseURL + "/datacenters/" + accountAlias

	dataCenters := model.DataCenterList{}
	a.GetWithAuthentication(endpoint, &dataCenters)

	return dataCenters
}
