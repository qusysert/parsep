package service

import (
	"parser/internal/model"
	"sync"
	"time"
)

func (s *Service) FormTableContent(tableData model.TableDataLinks) (model.TabledData, error) {
	var columns []model.TableColumn
	var wgExc sync.WaitGroup
	for _, pp := range tableData.ParsePool {
		wgExc.Add(1)
		go func(pp model.ParsePoint) {
			defer wgExc.Done()
			var col model.TableColumn
			col.Title = pp.Exchange + ", " + pp.PriceType + ", " + pp.Unit
			var wgUrl sync.WaitGroup
			for _, u := range pp.Urls {
				wgUrl.Add(1)
				go func(u model.Url) {
					defer wgUrl.Done()
					resp, _ := s.repo.ScrapePrice(u, pp.Formatter)
					col.Prices = append(col.Prices, resp)
				}(u)
			}
			wgUrl.Wait()
			columns = append(columns, col)
		}(pp)
	}
	now := time.Now()
	wgExc.Wait()
	return model.TabledData{TableTitle: tableData.Title + now.Format("02.01.06"), Columns: columns}, nil
}
