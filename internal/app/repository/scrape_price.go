package repository

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"parser/internal/model"
	"regexp"
	"strconv"
)

func (r *Repository) ScrapePrice(url model.Url, selectors model.ParseSelectors, formatter func(model.PriceRecord) model.PriceRecord) (model.PriceRecord, error) {
	var pr model.PriceRecord
	r.colly.OnHTML(selectors.PriceSelector, func(e *colly.HTMLElement) {
		price, err := strconv.ParseFloat(regexp.MustCompile(`[^\d\.]`).ReplaceAllString(e.Text, ""), 64)
		if err != nil {
			log.Println(err)
		}
		pr.Price = price
	})

	r.colly.OnHTML(selectors.ChangeSelector, func(e *colly.HTMLElement) {
		pr.Change = e.Text
	})

	err := r.colly.Visit(url.Link)
	if err != nil {
		return model.PriceRecord{}, fmt.Errorf("error visiting %s: %w", url, err)
	}
	pr.Material = url.Material
	pr = formatter(pr)
	return pr, nil
}
