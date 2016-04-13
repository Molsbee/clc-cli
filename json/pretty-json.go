package json

import (
	"encoding/json"
	"log"
)

// PrettyPrint return multiline formated json string
func PrettyPrint(body interface{}) string {
	json, err := json.MarshalIndent(body, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	return string(json)
}
