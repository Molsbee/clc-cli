package clc

import (
	"encoding/json"
	"github.com/molsbee/clc-cli/model"
	"github.com/prometheus/common/log"
	"net/http"
)

type API struct {
	BaseURL string
	BearerTokenAuthenticatedAPI
}

func NewAPI(authentication model.AuthenticationResponse) *API {
	return &API{
		BaseURL: "https://api.ctl.io/v2",
		BearerTokenAuthenticatedAPI: BearerTokenAuthenticatedAPI{
			Authentication: authentication,
			Client:         &http.Client{},
		},
	}
}

type BearerTokenAuthenticatedAPI struct {
	Authentication model.AuthenticationResponse
	Client         *http.Client
}

func (b *BearerTokenAuthenticatedAPI) GetWithAuthentication(endpoint string, resp interface{}) {
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+b.Authentication.BearerToken)

	response, err := b.Client.Do(req)
	if err != nil {
		log.Fatalf("error received while performing clc api request %s %v", endpoint, err)
	}

	json.NewDecoder(response.Body).Decode(resp)
}

func (b *BearerTokenAuthenticatedAPI) GetOrDefaultAccountAlias(alias string) string {
	if alias == "" {
		return b.Authentication.AccountAlias
	}
	return alias
}
