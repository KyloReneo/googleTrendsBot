package main

import (
	"fmt"

	"github.com/KyloReneo/googleTrendsBot/utils"

)

func main() {
	var country string = "US"

	data := utils.ReadGoogleTrends(country)
	rss := utils.XMLUnmarshal(data)

	fmt.Printf("Google search trends in %s for today:\n", country)
	for i := range rss.Channel.Items {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Printf("Search Term: %v in %s.\n", rss.Channel.Title, country)
		fmt.Println("Link to the Trend: ", rss.Channel.Items[i].Link)
		fmt.Println("	News items:")
		fmt.Println("	-----------------------------------------------------------------------------------------------------------------------------------")
		for j := range rss.Channel.Items[i].NewsItems {
			num := (j + 1)
			fmt.Printf("    	%d:\n", num)
			fmt.Println("	News item Title: ", rss.Channel.Items[i].NewsItems[j].NewsItemTitle)
			fmt.Println("	News item Link: ", rss.Channel.Items[i].NewsItems[j].NewsItemLink)
			fmt.Println("	-----------------------------------------------------------------------------------------------------------------------------------")
		}
	}

}
