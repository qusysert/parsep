package repository

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"log"
	"parser/internal/model"
	"regexp"
	"strconv"
)

func (r *Repository) ScrapePrice(url string, c colly.Collector, selectors model.ParseSelectors, formatter func(model.PriceRecord) model.PriceRecord) (model.PriceRecord, error) {
	var pr model.PriceRecord

	c.OnHTML(selectors.PriceSelector, func(e *colly.HTMLElement) {
		price, err := strconv.ParseFloat(regexp.MustCompile(`[^\d\.]`).ReplaceAllString(e.Text, ""), 64)
		if err != nil {
			log.Println(err)
		}
		pr.Price = price
	})

	c.OnHTML(selectors.TitleSelector, func(e *colly.HTMLElement) {
		commodity := e.DOM.Contents().FilterFunction(func(i int, s *goquery.Selection) bool {
			return s.Nodes[0].Type == html.TextNode
		}).Text()
		pr.Material = commodity
	})

	err := c.Visit(url)
	if err != nil {
		return model.PriceRecord{}, fmt.Errorf("error visiting %s: %w", url, err)
	}
	pr = formatter(pr)
	return pr, nil
}
