package service

import (
	"github.com/gocolly/colly"
	"parser/internal/model"
)

type Service struct {
	repo IRepository
}

// IRepository Interface for basic parsing
type IRepository interface {
	ScrapePrice(url string, c colly.Collector, selectors model.ParseSelectors, formatter func(model.PriceRecord) model.PriceRecord) (model.PriceRecord, error)
}

func New(r IRepository) *Service {
	return &Service{r}
}
