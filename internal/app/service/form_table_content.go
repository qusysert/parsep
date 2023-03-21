package service

import (
	"fmt"
	"parser/internal/model"
	"sort"
	"time"
)

func (s *Service) FormTableContent(tableData model.TableDataLinks) (model.TabledData, error) {
	var columns []model.TableColumn
	type result struct {
		resp model.PriceRecord
		err  error
	}
	for _, pp := range tableData.ParsePool {
		var col model.TableColumn
		results := make(chan result, len(pp.Urls))

		col.Title = pp.Exchange + ", " + pp.PriceType + ", " + pp.Unit
		for _, u := range pp.Urls {
			go func(url string) {
				for _, selector := range pp.Selectors {
					resp, err := s.repo.ScrapePrice(url, selector, pp.Formatter)
					results <- result{resp, err}
				}

			}(u)
		}
		for i := 0; i < len(pp.Urls); i++ {
			res := <-results
			if res.err != nil {
				return model.TabledData{}, fmt.Errorf("error pasing %s: %w", pp.Urls[i], res.err)
			}
			col.Prices = append(col.Prices, res.resp)
		}

		sort.Slice(col.Prices, func(i, j int) bool {
			return col.Prices[i].Material < col.Prices[j].Material
		})

		columns = append(columns, col)
	}
	now := time.Now()

	return model.TabledData{TableTitle: tableData.Title + now.Format("02.01.06"), Columns: columns}, nil
}
