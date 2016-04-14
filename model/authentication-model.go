package model

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

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

// DecodeBearerToken Returns a decoded bearer token details
func (a AuthenticationResponse) DecodeBearerToken() (*BearerTokenDetails, error) {
	if a.BearerToken == "" {
		return nil, fmt.Errorf("PreAuthentication failed please reauthenticate")
	}

	bearerToken := strings.Split(a.BearerToken, ".")[1]
	if l := len(bearerToken) % 4; l > 0 {
		bearerToken += strings.Repeat("=", 4-l)
	}

	resp, err := base64.URLEncoding.DecodeString(bearerToken)
	if err != nil {
		log.Fatalln("Failed to Decode BearerTokenDetails", err)
	}

	tokenDetails := BearerTokenDetails{}
	if err := json.Unmarshal(resp, &tokenDetails); err != nil {
		log.Fatalln("Unable to unmarshal bearerToken to BearerTokenDetails", err)
	}

	return &tokenDetails, nil
}

// BearerTokenDetails Decode information from Second Section of BearerToken
type BearerTokenDetails struct {
	ISS             string `json:"iss"`
	AUD             string `json:"aud"`
	NBF             int64  `json:"nbf"`
	EXP             int64  `json:"exp"`
	UniqueName      string `json:"unique_name"`
	AccountAlias    string `json:"urn:tier3:account-alias"`
	DefaultLocation string `json:"urn:tier3:location-alias"`
	Role            string `json:"role"`
}
