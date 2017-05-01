package api

import "github.com/molsbee/clc-cli/authentication"

type BaseRequest struct {
	AccountAlias string
}

func (b BaseRequest) GetAccountAlias() string {
	if b.AccountAlias == "" {
		return authentication.Auth.AccountAlias
	}

	return b.AccountAlias
}