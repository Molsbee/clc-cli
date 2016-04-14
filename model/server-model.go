package model

// Details Response from CLC V2 API Get Server
type Details struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	GroupID     string `json:"groupId"`
	LocationID  string `json:"locationId"`
	OsType      string `json:"osType"`
	Status      string `json:"status"`
	Details     struct {
		IPAddresses []struct {
			Public   string `json:"public"`
			Internal string `json:"internal"`
		} `json:"ipAddresses"`
	} `json:"details"`
	Type        string `json:"type"`
	StorageType string `json:"storageType"`
}

// Credentials Response from CLC V2 API Get Server Credentials
type Credentials struct {
	Username string `json:"userName"`
	Password string `json:"password"`
}
