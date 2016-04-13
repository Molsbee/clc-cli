package model

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

// DataCenterList List of DataCenters
type DataCenterList []struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
