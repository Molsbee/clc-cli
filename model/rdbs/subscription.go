package rdbs

import (
	"bytes"
	"fmt"
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

func (s Subscription) String() string {
	return fmt.Sprint(s.FmtString(""))
}

func (s Subscription) FmtString(format string) string {
	buf := new(bytes.Buffer)
	buf.WriteString(format + "ID:\t" + strconv.Itoa(s.ID) + "\n")
	buf.WriteString(format + "AccountAlias:\t" + s.AccountAlias + "\n")
	buf.WriteString(format + "Location:\t" + s.Location + "\n")
	buf.WriteString(format + "InstanceType:\t" + s.InstanceType + "\n")
	buf.WriteString(format + "Engine:\t" + s.Engine + "\n")
	if s.Edition != "" {
		buf.WriteString(format + "Edition:\t" + s.Edition + "\n")
	}
	if s.Version != "" {
		buf.WriteString(format + "Version:\t" + s.Version + "\n")
	}
	buf.WriteString(format + "External ID:\t" + s.ExternalID + "\n")
	buf.WriteString(format + "Restart Required:\t" + strconv.FormatBool(s.RestartRequired) + "\n")
	buf.WriteString(format + "BackupTime:\t" + s.BackupTime + "\n")
	buf.WriteString(format + "Backup Retention Days:\t" + strconv.Itoa(s.BackupRetentionDays) + "\n")
	buf.WriteString(format + "Servers:\n")
	for _, server := range s.Servers {
		buf.WriteString(server.FmtString("\t"))
	}
	return fmt.Sprint(buf.String())
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

func (se Server) String() string {
	return fmt.Sprint(se.FmtString(""))
}

func (se Server) FmtString(format string) string {
	buf := new(bytes.Buffer)
	buf.WriteString(format + "ID:\t" + strconv.Itoa(se.ID) + "\n")
	buf.WriteString(format + "Alias:\t" + se.Alias + "\n")
	buf.WriteString(format + "Location:\t" + se.Location + "\n")
	buf.WriteString(format + "CPU:\t" + strconv.Itoa(se.CPU) + "\n")
	buf.WriteString(format + "Memory:\t" + strconv.Itoa(se.Memory) + "\n")
	buf.WriteString(format + "Storage:\t" + strconv.Itoa(se.Storage) + "\n")
	buf.WriteString(format + "Attributes:\n")
	for key, value := range se.Attributes {
		buf.WriteString(fmt.Sprintf(format + "\t%s:\t%s\n", key, value))
	}
	buf.WriteString(format + "Connections:\t" + strconv.Itoa(se.Connections) + "\n")
	return fmt.Sprint(buf.String())
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
