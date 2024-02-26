package collector

import (
	"encoding/json"
	"log"
)

func receive(message string) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal([]byte(message), &data)
	if err != nil {
		log.Println("Failed to unmarshal message:", err)
	}
	return data, err
}
