package parser

import (
	"github.com/gocolly/colly"
	"log"
	"time"
)

type Parser struct {
	Collector colly.Collector
}

func New(timeout int) *Parser {
	collector := colly.NewCollector(
		colly.AllowURLRevisit(),
	)
	collector.SetRequestTimeout(time.Duration(timeout) * time.Second)
	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		log.Println("Got a response from", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, e error) {
		log.Println("Got this error:", e)
	})
	return &Parser{
		Collector: *collector,
	}
}
