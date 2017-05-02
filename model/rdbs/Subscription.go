package rdbs

import (
	"bytes"
	"strconv"
)

type Subscription struct {
	ID                   int                  `json:"id"`
	AccountAlias         string               `json:"accountAlias"`
	Location             string               `json:"location"`
	InstanceType         string               `json:"instanceType"`
	Engine               string               `json:"engine"`
	Edition              string               `json:"edition"`
	Version              string               `json:"version"`
	ExternalID           string               `json:"externalId"`
	RestartRequired      bool                 `json:"restartRequired"`
	Status               string               `json:"status"`
	BackupTime           string               `json:"backupTime"`
	BackupRetentionDays  int                  `json:"backupRetentionDays"`
	Users                []User               `json:"users"`
	Instances            []Instance           `json:"instances"`
	Servers              []Server             `json:"servers"`
	Host                 string               `json:"host"`
	Port                 int                  `json:"port"`
	Certificate          string               `json:"certificate"`
	Backups              []Backup             `json:"backups"`
	ConfigurationProfile ConfigurationProfile `json:"configurationProfile"`
	PendingJobCount      int                  `json:"pendingJobCount"`
}

func (s Subscription) ToString() {
	buf := new(bytes.Buffer)
	buf.WriteString("ID: " + strconv.Itoa(s.ID) + "/n")
	buf.WriteString("AccountAlias: " + s.AccountAlias + "/n")

}

type SubscriptionReference struct {
	ExternalID string `json:"externalId"`
	ID         int    `json:"id"`
}

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Instance struct {
	Name string `json:"name"`
}

type Server struct {
	ID          int               `json:"id"`
	Alias       string            `json:"alias"`
	Location    string            `json:"location"`
	CPU         int               `json:"cpu"`
	Memory      int               `json:"memory"`
	Storage     int               `json:"storage"`
	Attributes  map[string]string `json:"attributes"`
	Connections int               `json:"connections"`
}

type Backup struct {
	ID         int    `json:"id"`
	FileName   string `json:"fileName"`
	BackupTime string `json:"backupTime"`
	BackupType string `json:"backupType"`
	Status     string `json:"status"`
	Size       int    `json:"size"`
}

type ConfigurationProfile struct {
	ID            int                     `json:"id"`
	Name          string                  `json:"name"`
	Description   string                  `json:"description"`
	LastEditedBy  string                  `json:"lastEditedBy"`
	LastEdited    string                  `json:"lastEdited"`
	Parameters    map[string]string       `json:"parameters"`
	IsDefault     bool                    `json:"isDefault"`
	Subscriptions []SubscriptionReference `json:"subscriptions"`
}
