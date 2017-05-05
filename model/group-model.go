package model

import (
	"fmt"
	"time"
	"strings"
)

const server = "server"

// Group Model to Represent data being returned for a Particular Group
type Group struct {
	ID           string  `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	LocationID   string  `json:"locationId"`
	Type         string  `json:"type"`
	Status       string  `json:"status"`
	ServersCount int     `json:"serversCount"`
	Groups       []Group `json:"groups"`
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
}

// GetServers Returns formatted string of server aliases for a given group
func (g Group) GetServers() string {
	var resp string
	for _, link := range g.Links {
		if link.Rel == server {
			resp += fmt.Sprintln("ID: ", link.ID)
		}
	}

	for _, group := range g.Groups {
		resp += group.GetServers()
	}

	return resp
}

func (g Group) String() string {
	return g.fmtString("")
}

func (g Group) fmtString(s string) string {
	resp := fmt.Sprintf("%sName:\t\t%s\n", s, g.Name)
	resp += fmt.Sprintf("%sID:\t\t%s\n", s, g.ID)
	resp += fmt.Sprintf("%sServer Count:\t%d\n", s, g.ServersCount)

	for pos, link := range g.Links {
		if link.Rel == server {
			resp += fmt.Sprintf("%s\t%s", s, link.ID)
			if pos % 5 == 0 || pos == len(g.Links) - 1 {
				resp += fmt.Sprintln()
			}
		}
	}

	for _, group := range g.Groups {
		if strings.ToLower(group.Name) == "trash" {
			continue
		}

		resp += fmt.Sprintf("%s", group.fmtString(s+"\t"))
	}

	return resp
}
