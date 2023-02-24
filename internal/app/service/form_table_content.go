package service

import (
	"fmt"
	"github.com/gocolly/colly"
	"parser/internal/model"
)

func (s *Service) FormTableContent(c colly.Collector, pool []model.ParsePoint) ([]model.TableColumn, error) {
	var columns []model.TableColumn
	type result struct {
		resp model.PriceRecord
		err  error
	}

	for _, pp := range pool {
		var col model.TableColumn
		results := make(chan result, len(pp.Urls))

		col.Title = pp.Exchange + ", " + pp.PriceType + ", " + pp.Unit
		for _, u := range pp.Urls {
			go func(url string) {
				resp, err := s.repo.ScrapePrice(url, c, pp.Selectors, pp.Formatter)

				results <- result{resp, err}

			}(u)
		}
		for i := 0; i < len(pp.Urls); i++ {
			res := <-results
			if res.err != nil {
				return nil, fmt.Errorf("error pasing %s: %w", pp.Urls[i], res.err)
			}
			col.Prices = append(col.Prices, res.resp)
		}
		columns = append(columns, col)

	}

	return columns, nil
}
