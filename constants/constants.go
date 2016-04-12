package constants

import "net/http"

var (
	// ClcAPIEndpoint Base Endpoint for CLC API
	ClcAPIEndpoint = "https://api.ctl.io/v2"

	// ClcAuthenticaitonEndpoint Authenticaiton Endpoint
	ClcAuthenticaitonEndpoint = ClcAPIEndpoint + "/authentication/login"

	// Client Global Http Client
	Client = &http.Client{}
)
