package main

import (
	"fmt"
	"log"

	"github.com/KyloReneo/googleTrendsBot/utils"
)

func main() {

	var country string

	fmt.Println("Specify the country that you want to see it's google trends by entering it's name.")
	fmt.Print("for example US for Unitedstates: ")
	_, err := fmt.Scanln(&country)
	if err != nil {
		log.Fatalf("Reading the input went wrong due to the original error:\n %v", err)
	}

	data := utils.ReadGoogleTrends(country)
	rss := utils.XMLUnmarshal(data)

	jsonRSS := utils.JSONRSS{
		RSSName: "google trends rss",
		Channels: []utils.JSONChannel{
			{
				Title: rss.Channel.Title,
				Items: []utils.JSONItem{},
			},
		},
	}

	fmt.Println("---------------------------------------------------------------------------------------------------------------------------------------")
	fmt.Printf("Google search trends in %s for today:\n", country)
	for i := range rss.Channel.Items {

		rank := (i + 1)

		newItem := utils.JSONItem{
			ItemRank:  rank,
			Title:     rss.Channel.Items[i].Title,
			Link:      rss.Channel.Items[i].Link,
			Traffic:   rss.Channel.Items[i].Traffic,
			NewsItems: []utils.JSONNewsItems{},
		}

		jsonRSS.Channels[0].Items = append(jsonRSS.Channels[0].Items, newItem)

		fmt.Println("#", rank)
		fmt.Printf("Search Term: %v in %s.\n", rss.Channel.Title, country)
		fmt.Println("Link to the Trend: ", rss.Channel.Items[i].Link)
		fmt.Println("	News items:")
		fmt.Println("	-----------------------------------------------------------------------------------------------------------------------------------")
		for j := range rss.Channel.Items[i].NewsItems {

			num := (j + 1)

			newNewsItem := utils.JSONNewsItems{
				NewsItemsNumber: num,
				NewsItemTitle:   rss.Channel.Items[i].NewsItems[j].NewsItemTitle,
				NewsItemLink:    rss.Channel.Items[i].NewsItems[j].NewsItemLink,
			}

			jsonRSS.Channels[0].Items[i].NewsItems = append(jsonRSS.Channels[0].Items[i].NewsItems, newNewsItem)

			fmt.Printf("    	%d:\n", num)
			fmt.Println("	News item Title: ", rss.Channel.Items[i].NewsItems[j].NewsItemTitle)
			fmt.Println("	News item Link: ", rss.Channel.Items[i].NewsItems[j].NewsItemLink)
			fmt.Println("	-----------------------------------------------------------------------------------------------------------------------------------")
		}
	}

	jsonStr := utils.JSONMarshal(jsonRSS)

	utils.WriteToJSON(jsonStr)
}
