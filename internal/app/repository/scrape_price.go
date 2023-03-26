package repository

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"parser/internal/app/pkg/utils"
	"parser/internal/model"
)

func (r *Repository) ScrapePrice(url model.Url, formatter func(model.PriceRecord) model.PriceRecord) (model.PriceRecord, error) {
	var pr model.PriceRecord

	r.colly.OnHTML(url.Selectors.PriceSelector, func(e *colly.HTMLElement) {
		price, err := utils.FormatPrice(e.Text)
		if err != nil {
			log.Println("Error parsing price: " + err.Error())
			return
		}
		pr.Price = price
	})

	if url.Selectors.ChangeSelector != "" {
		r.colly.OnHTML(url.Selectors.ChangeSelector, func(e *colly.HTMLElement) {
			pr.Change = e.Text
		})
	} else if url.Selectors.PrevPriceSelector != "" {
		var prevPrice float64
		r.colly.OnHTML(url.Selectors.PrevPriceSelector, func(e *colly.HTMLElement) {
			var err error
			prevPrice, err = utils.FormatPrice(e.Text)
			if err != nil {
				log.Println("Error parsing prevPrice: " + err.Error())
			}
			pr.Change = fmt.Sprintf("%.2f", ((pr.Price-prevPrice)/prevPrice)*100)
		})

	}

	err := r.colly.Visit(url.Link)
	if err != nil {
		return model.PriceRecord{}, fmt.Errorf("error visiting %s: %w", url, err)
	}
	pr.Material = url.Material
	pr = formatter(pr)
	return pr, nil
}
