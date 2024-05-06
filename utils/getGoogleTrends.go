package utils

import (
	"fmt"
	"log"
	"net/http"
)

func GetGoogleTrends(country string) *http.Response {
	rssXML, err := http.Get(fmt.Sprintf("https://trends.google.com/trends/trendingsearches/daily/rss?geo=%s", country))
	if err != nil {
		log.Fatalf("Requesting the RSS feed went wrong, due to the original error:\n %v", err)
	}
	return rssXML
}
