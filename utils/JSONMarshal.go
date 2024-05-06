package utils

import (
	"encoding/json"
	"log"

)

func JSONMarshal(data JSONRSS) string {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Marshalling json data went wrong due to the original error:\n %v", err)
	}
	return string(bytes)
}
