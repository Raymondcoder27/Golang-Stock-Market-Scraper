package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func main() {
	ticker := []string{
		"MSFT",
		"IBM",
		"GE",
		"UNP",
		"COST",
		"MCD",
		"V",
		"WMT",
		"MMM",
		"AXP",
		"AAPL",
		"BA",
		"CSCO",
		"GS",
		"JPM",
		"CRM",
		"VZ",
	}

	stocks := []Stock{}

	c := colly.NewCollector()
	c.OnRequest(func(r *colly.Request) {
		fmt.Print("Visiting", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong", err)
	})

}
