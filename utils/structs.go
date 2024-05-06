package utils

import "encoding/xml"

type RSS struct {
	XMLName xml.Name
	Channel *Channel
}

type Channel struct {
	Title string
	Items []Item
}

type Item struct {
	Title     string
	Link      string
	Traffic   string
	NewsItems []NewsItems
}

type NewsItems struct {
	NewsItemTitle string
	NewsItemLink  string
}
