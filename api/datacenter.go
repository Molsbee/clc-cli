package api

import (
	"github.com/molsbee/clc-cli/authentication"
	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model"
	"github.com/molsbee/clc-cli/net"
)

// DataCenter represents name of datacenter
type DataCenter struct {
	Name string
}

// Get Returns Data Center details
func (d DataCenter) Get() model.DataCenter {
	endpoint := constants.ClcAPIEndpoint + "/datacenters/" + authentication.Auth.AccountAlias + "/" + d.Name + "?groupLinks=true"

	dataCenter := model.DataCenter{}
	net.GetWithAuthentication(endpoint, &dataCenter)

	return dataCenter
}

// List Returns List of Data Centers
func (d DataCenter) List() model.DataCenterList {
	endpoint := constants.ClcAPIEndpoint + "/datacenters/" + authentication.Auth.AccountAlias

	dataCenters := model.DataCenterList{}
	net.GetWithAuthentication(endpoint, &dataCenters)

	return dataCenters
}
