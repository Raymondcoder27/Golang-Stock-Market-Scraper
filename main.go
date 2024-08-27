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

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}
		stock.company = e.ChildText("h1")
		fmt.Println("Company", stock.company)
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		fmt.Println("Pice", stock.price)
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")
		fmt.Println("Change", stock.change)

		stocks = append(stocks, stock)
	})

	c.Visit("https://finance.yahoo.com/qutote/" + t + "/")
}
