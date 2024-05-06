package utils

import (
	"io"
	"log"
)

func ReadGoogleTrends(country string) []byte {
	resp := GetGoogleTrends(country)
	rssData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Reading the RSS data went wrong. Original error:\n %v", err)
	}
	return rssData
}
