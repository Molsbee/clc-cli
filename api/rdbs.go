package api

import (
	"github.com/molsbee/clc-cli/constants"
	"github.com/molsbee/clc-cli/model/rdbs"
	"github.com/molsbee/clc-cli/net"
)

type RdbsSubscription struct {
	BaseRequest
	SubscriptionId string
}

func (r RdbsSubscription) Get() interface{} {
	endpoint := getBaseEndpoint(r.GetAccountAlias())
	subscriptionId := r.SubscriptionId

	endpoint += "/" + subscriptionId
	subscription := rdbs.Subscription{}
	net.GetWithAuthentication(endpoint, &subscription)
	return subscription
}

func (r RdbsSubscription) GetAll() []rdbs.Subscription {
	endpoint := getBaseEndpoint(r.GetAccountAlias())

	subscriptions := []rdbs.Subscription{}
	net.GetWithAuthentication(endpoint, &subscriptions)

	return subscriptions
}

func getBaseEndpoint(accountAlias string) string {
	return constants.RDBS["PROD"] + "/v1/" + accountAlias + "/subscriptions"
}
