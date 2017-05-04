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
	Edition              string               `json:"edition,omitempty"`
	Version              string               `json:"version,omitempty"`
	ExternalID           string               `json:"externalId"`
	RestartRequired      bool                 `json:"restartRequired"`
	Status               string               `json:"status"`
	BackupTime           string               `json:"backupTime"`
	BackupRetentionDays  int                  `json:"backupRetentionDays"`
	Users                []User               `json:"users,omitempty"`
	Instances            []Instance           `json:"instances,omitempty"`
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
	buf.WriteString(format + "ID:\t\t\t" + strconv.Itoa(s.ID) + "\n")
	buf.WriteString(format + "AccountAlias:\t\t" + s.AccountAlias + "\n")
	buf.WriteString(format + "External ID:\t\t" + s.ExternalID + "\n")
	buf.WriteString(format + "Location:\t\t" + s.Location + "\n")
	buf.WriteString(format + "Engine:\t\t\t" + s.Engine + "\n")
	if s.Edition != "" {
		buf.WriteString(format + "Edition:\t\t" + s.Edition + "\n")
	}
	if s.Version != "" {
		buf.WriteString(format + "Version:\t\t" + s.Version + "\n")
	}
	buf.WriteString(format + "Status:\t\t\t" + s.Status + "\n")
	buf.WriteString(format + "BackupTime:\t\t" + s.BackupTime + "\n")
	buf.WriteString(format + "Backup Retention Days:\t" + strconv.Itoa(s.BackupRetentionDays) + "\n")
	buf.WriteString(format + "Connection:\t\t" + s.Host + ":" + strconv.Itoa(s.Port) + "\n")
	buf.WriteString(format + "Servers:\n")
	for _, server := range s.Servers {
		buf.WriteString(server.FmtString(format + "\t"))
	}
	buf.WriteString(format + "Backups:\n")
	for _, backup := range s.Backups {
		buf.WriteString(backup.FmtString(format + "\t"))
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
	internalIP := "UNKNOWN"
	for key, value := range se.Attributes {
		if key == "INTERNAL_IP" {
			internalIP = value
		}
	}

	return fmt.Sprintf(format+"Alias: %s IP: %s\n", se.Alias, internalIP)
}

type Backup struct {
	ID         int    `json:"id"`
	FileName   string `json:"fileName"`
	BackupTime int    `json:"backupTime"`
	BackupType string `json:"backupType"`
	Status     string `json:"status"`
	Size       int    `json:"size"`
}

func (b Backup) String() string {
	return b.FmtString("")
}

func (b Backup) FmtString(format string) string {
	return fmt.Sprintf(format+"FileName: %s Status: %s\n", b.FileName, b.Status)
}

type ConfigurationProfile struct {
	ID            int                     `json:"id"`
	Name          string                  `json:"name"`
	Description   string                  `json:"description"`
	LastEditedBy  string                  `json:"lastEditedBy,omitempty"`
	LastEdited    int                     `json:"lastEdited"`
	Parameters    map[string]string       `json:"parameters"`
	IsDefault     bool                    `json:"isDefault"`
	Subscriptions []SubscriptionReference `json:"subscriptions"`
}
