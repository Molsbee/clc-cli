package model

// AuthenticationRequest Requests to authenticate with CLC V2 API
type AuthenticationRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// AuthenticationResponse Response from CLC V2 API Authenticaiton Service
type AuthenticationResponse struct {
	Username      string `json:"userName"`
	AccountAlias  string `json:"accountAlias"`
	LocationAlias string `json:"locationAlias"`
	BearerToken   string `json:"bearerToken"`
}
