package service

import (
	"parser/internal/model"
	"time"
)

func (s *Service) FormTableContent(tableData model.TableDataLinks) (model.TabledData, error) {
	var columns []model.TableColumn
	for _, pp := range tableData.ParsePool {
		var col model.TableColumn

		col.Title = pp.Exchange + ", " + pp.PriceType + ", " + pp.Unit
		for _, u := range pp.Urls {
			resp, _ := s.repo.ScrapePrice(u, pp.Formatter)
			col.Prices = append(col.Prices, resp)
		}

		columns = append(columns, col)
	}
	now := time.Now()

	return model.TabledData{TableTitle: tableData.Title + now.Format("02.01.06"), Columns: columns}, nil
}
