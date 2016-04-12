package model

import "time"

// DataCenter Model to Represent data being returned by DataCenter Api
type DataCenter struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Links []struct {
		Rel   string   `json:"rel"`
		Href  string   `json:"href"`
		Verbs []string `json:"verbs,omitempty"`
		ID    string   `json:"id,omitempty"`
		Name  string   `json:"name,omitempty"`
	} `json:"links"`
}

// Group Model to Represent data being returned for a Particular Group
type Group struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	LocationID   string `json:"locationId"`
	Type         string `json:"type"`
	Status       string `json:"status"`
	ServersCount int    `json:"serversCount"`
	Groups       []struct {
		ID           string        `json:"id"`
		Name         string        `json:"name"`
		Description  string        `json:"description,omitempty"`
		LocationID   string        `json:"locationId"`
		Type         string        `json:"type"`
		Status       string        `json:"status"`
		ServersCount int           `json:"serversCount"`
		Groups       []interface{} `json:"groups"`
		Links        []struct {
			Rel   string   `json:"rel"`
			Href  string   `json:"href"`
			Verbs []string `json:"verbs,omitempty"`
			ID    string   `json:"id,omitempty"`
		} `json:"links"`
		ChangeInfo struct {
			CreatedBy    string    `json:"createdBy"`
			CreatedDate  time.Time `json:"createdDate"`
			ModifiedBy   string    `json:"modifiedBy"`
			ModifiedDate time.Time `json:"modifiedDate"`
		} `json:"changeInfo"`
		CustomFields []interface{} `json:"customFields"`
	} `json:"groups"`
	Links []struct {
		Rel   string   `json:"rel"`
		Href  string   `json:"href"`
		Verbs []string `json:"verbs,omitempty"`
	} `json:"links"`
	ChangeInfo struct {
		CreatedBy    string    `json:"createdBy"`
		CreatedDate  time.Time `json:"createdDate"`
		ModifiedBy   string    `json:"modifiedBy"`
		ModifiedDate time.Time `json:"modifiedDate"`
	} `json:"changeInfo"`
	CustomFields []interface{} `json:"customFields"`
}
