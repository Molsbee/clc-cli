package api

import (
	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model"
	"github.com/molsbee/clc-cli/net"
)

// DataCenter struct represents the necessary data to perform DataCenter API request to CL V2 APId
type DataCenter struct {
	BaseRequest
	Name string
}

// Get Returns Data Center details
func (d DataCenter) Get() model.DataCenter {
	endpoint := constants.ClcAPIEndpoint + "/datacenters/" + d.GetAccountAlias() + "/" + d.Name + "?groupLinks=true"

	dataCenter := model.DataCenter{}
	net.GetWithAuthentication(endpoint, &dataCenter)

	return dataCenter
}

// List Returns List of Data Centers
func (d DataCenter) List() model.DataCenterList {
	endpoint := constants.ClcAPIEndpoint + "/datacenters/" + d.GetAccountAlias()

	dataCenters := model.DataCenterList{}
	net.GetWithAuthentication(endpoint, &dataCenters)

	return dataCenters
}
