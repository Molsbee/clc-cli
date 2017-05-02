package constants

import "net/http"

var (
	ClcAPIEndpoint = "https://api.ctl.io/v2"
	ClcAuthenticaitonEndpoint = ClcAPIEndpoint + "/authentication/login"

	RDBS = map[string]string{
		"DEV": "https://api-dv.rdbs.ctl.io",
		"QA": "https://api-qa.rdbs.ctl.io",
		"PROD": "https://api.rdbs.ctl.io",
	}

	Client = &http.Client{}
)
