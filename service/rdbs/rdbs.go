package rdbs

import (
	"github.com/molsbee/clc-cli/model"
	"github.com/molsbee/clc-cli/model/rdbs"
	"github.com/molsbee/clc-cli/service/clc"
	"net/http"
	"strconv"
)

var (
	RDBS = map[string]string{
		"DEV":  "https://api-dv.rdbs.ctl.io",
		"QA":   "https://api-qa.rdbs.ctl.io",
		"PROD": "https://api.rdbs.ctl.io",
	}
)

type API struct {
	clc.BearerTokenAuthenticatedAPI
}

func NewAPI(authentication model.AuthenticationResponse) *API {
	return &API{
		BearerTokenAuthenticatedAPI: clc.BearerTokenAuthenticatedAPI{
			Authentication: authentication,
			Client:         &http.Client{},
		},
	}
}

type EnvironmentRequest struct {
	Environment string
}

func (e EnvironmentRequest) getBaseURL() string {
	if e.Environment != "" {
		for key, value := range RDBS {
			if key == e.Environment {
				return value
			}
		}
	}

	return RDBS["PROD"]
}

type SubscriptionRequest struct {
	AccountAlias   string
	SubscriptionID int
	EnvironmentRequest
}

func (a *API) GetSubscriptions(request SubscriptionRequest) []rdbs.Subscription {
	accountAlias := a.GetOrDefaultAccountAlias(request.AccountAlias)
	endpoint := request.getBaseURL() + "/v1/" + accountAlias + "/subscriptions"

	subscriptions := []rdbs.Subscription{}
	a.GetWithAuthentication(endpoint, &subscriptions)

	return subscriptions
}

func (a *API) GetSubscription(request SubscriptionRequest) rdbs.Subscription {
	accountAlias := a.GetOrDefaultAccountAlias(request.AccountAlias)
	endpoint := request.getBaseURL() + "/v1/" + accountAlias + "/subscriptions/" + strconv.Itoa(request.SubscriptionID)

	subscription := rdbs.Subscription{}
	a.GetWithAuthentication(endpoint, &subscription)

	return subscription
}
