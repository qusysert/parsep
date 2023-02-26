package service

import (
	"github.com/gocolly/colly"
	"parser/internal/model"
	"parser/pkg/render_client"
)

type Service struct {
	repo   IRepository
	render IRender
}

// IRepository Interface for basic parsing
type IRepository interface {
	ScrapePrice(url string, c colly.Collector, selectors model.ParseSelectors, formatter func(model.PriceRecord) model.PriceRecord) (model.PriceRecord, error)
}

type IRender interface {
	Render(req render_client.Request) ([]byte, error)
}

func New(repo IRepository, render IRender) *Service {
	return &Service{repo, render}
}
