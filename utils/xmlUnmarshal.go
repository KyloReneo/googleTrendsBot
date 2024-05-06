package utils

import (
	"encoding/xml"
	"log"

)

func XMLUnmarshal(data []byte) RSS {
	var result RSS
	err := xml.Unmarshal(data, &result)
	if err != nil {
		log.Fatalf("Unmarshaling went wrong due to the original error:\n %v", err)
	}
	return result
}
