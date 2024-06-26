package utils

import "encoding/xml"

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title string `xml:"title"`
	Items []Item `xml:"item"`
}

type Item struct {
	Title     string      `xml:"title"`
	Link      string      `xml:"link"`
	Traffic   string      `xml:"approx_traffic"`
	NewsItems []NewsItems `xml:"news_item"`
}

type NewsItems struct {
	NewsItemTitle string `xml:"news_item_title"`
	NewsItemLink  string `xml:"news_item_url"`
}

type JSONRSS struct {
	RSSName  string
	Channels []JSONChannel
}

type JSONChannel struct {
	Title string
	Items []JSONItem
}

type JSONItem struct {
	ItemRank  int
	Title     string
	Link      string
	Traffic   string
	NewsItems []JSONNewsItems
}

type JSONNewsItems struct {
	NewsItemsNumber int
	NewsItemTitle   string
	NewsItemLink    string
}
