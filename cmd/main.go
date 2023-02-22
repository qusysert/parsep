package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	_ "github.com/gocolly/colly"
	"golang.org/x/net/html"
	"log"
	"parser/cmd/internal/model"
	"regexp"
	"strconv"
	"sync"
	"time"
)

func main() {
	var materials []model.PriceRecord
	mcxSelectors := model.ParseSelectors{
		TitleSelector: ".commodityTitle",
		PriceSelector: ".commodityPrice",
	}
	lmeSelectors := model.ParseSelectors{
		TitleSelector: ".hero-banner__title",
		PriceSelector: ".hero-metal-data__number",
	}
	start := time.Now()

	var wg sync.WaitGroup

	parsePool := []model.ParsePool{
		{
			Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-NICKEL.cms",
			Selectors: mcxSelectors,
			Formatter: model.McxFormatter,
		},
		{
			Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-ALUMINIUM.cms",
			Selectors: mcxSelectors,
			Formatter: model.McxFormatter,
		},
		{
			Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-ZINC.cms",
			Selectors: mcxSelectors,
			Formatter: model.McxFormatter,
		},
		{
			Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-LEAD.cms",
			Selectors: mcxSelectors,
			Formatter: model.McxFormatter,
		},
		{
			Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-COPPER.cms",
			Selectors: mcxSelectors,
			Formatter: model.McxFormatter,
		},
		{
			Url:       "https://www.lme.com/Metals/Non-ferrous/LME-Copper#Trading+day+summary",
			Selectors: lmeSelectors,
			Formatter: model.LmeFormatter,
		},
		{
			Url:       "https://www.lme.com/en/Metals/Non-ferrous/LME-Aluminium#Trading+day+summary",
			Selectors: lmeSelectors,
			Formatter: model.LmeFormatter,
		},
		{
			Url:       "https://www.lme.com/en/Metals/Non-ferrous/LME-Nickel",
			Selectors: lmeSelectors,
			Formatter: model.LmeFormatter,
		},
		{
			Url:       "https://www.lme.com/en/Metals/Non-ferrous/LME-Zinc#Trading+day+summary",
			Selectors: lmeSelectors,
			Formatter: model.LmeFormatter,
		},
	}

	collector := colly.NewCollector()
	collector.SetRequestTimeout(120 * time.Second)
	collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	collector.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	collector.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})

	// Scrape each URL concurrently using a goroutine
	for _, parsePoint := range parsePool {
		wg.Add(1)
		go func(u string, s model.ParseSelectors, f func(model.PriceRecord) model.PriceRecord) {
			defer wg.Done()
			pr, err := ScrapePrice(u, *collector, s, f)
			if err != nil {
				log.Printf("Error scraping %s: %v\n", u, err)
				return
			}
			materials = append(materials, pr)
		}(parsePoint.Url, parsePoint.Selectors, parsePoint.Formatter)
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
		log.Println("scraping price")
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
