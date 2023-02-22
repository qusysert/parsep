package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	_ "github.com/gocolly/colly"
	"golang.org/x/net/html"
	"log"
	model "parser/internal/model"
	"regexp"
	"strconv"
	"sync"
	"time"
)

func main() {
	var materials []model.PriceRecord
	start := time.Now()
	var wg sync.WaitGroup

	collector := colly.NewCollector()
	collector.SetRequestTimeout(120 * time.Second)
	collector.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		log.Println("Got a response from", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, e error) {
		log.Println("Got this error:", e)
	})

	// Scrape each URL concurrently using a goroutine
	for _, parsePoint := range model.ParsePool {
		wg.Add(1)
		go func(u string, s model.ParseSelectors, f func(model.PriceRecord) model.PriceRecord, e string) {
			defer wg.Done()
			pr, err := ScrapePrice(u, *collector, s, f)
			pr.Exchange = e
			if err != nil {
				log.Printf("Error scraping %s: %v\n", u, err)
				return
			}
			materials = append(materials, pr)
		}(parsePoint.Url, parsePoint.Selectors, parsePoint.Formatter, parsePoint.Exchange)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println(materials)
	fmt.Printf("Scraped in %v\n", time.Since(start))
}

func ScrapePrice(url string, c colly.Collector, selectors model.ParseSelectors, formatter func(model.PriceRecord) model.PriceRecord) (model.PriceRecord, error) {
	pr := model.PriceRecord{}

	c.OnHTML(selectors.TitleSelector, func(e *colly.HTMLElement) {
		commodity := e.DOM.Contents().FilterFunction(func(i int, s *goquery.Selection) bool {
			return s.Nodes[0].Type == html.TextNode
		}).Text()
		pr.Material = commodity
	})

	c.OnHTML(selectors.PriceSelector, func(e *colly.HTMLElement) {
		price, err := strconv.ParseFloat(regexp.MustCompile(`[^\d\.]`).ReplaceAllString(e.Text, ""), 64)
		if err != nil {
			log.Println(err)
		}
		pr.Price = price
	})

	err := c.Visit(url)
	if err != nil {
		return model.PriceRecord{}, fmt.Errorf("error visiting %s: %w", url, err)
	}
	pr = formatter(pr)
	return pr, nil
}
