package model

import "fmt"

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

func (d DataCenter) String() string {
	resp := fmt.Sprintf("ID:\t%s\n", d.ID)
	resp += fmt.Sprintf("Name:\t%s\n", d.Name)
	for _, element := range d.Links {
		if element.Rel == "group" {
			resp += fmt.Sprintln("Group")
			resp += fmt.Sprintf("\tName:\t%s\n", element.Name)
			resp += fmt.Sprintf("\tID:\t%s\n", element.ID)
		}
	}

	return resp
}

// DataCenterList List of DataCenters
type DataCenterList []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (l DataCenterList) String() string {
	var resp string
	for _, element := range l {
		resp += fmt.Sprintf("ID: %s\t\tName: %s\n", element.ID, element.Name)
	}

	return resp
}
