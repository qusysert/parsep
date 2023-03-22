package service

import (
	"parser/internal/model"
)

type Service struct {
	repo   IRepository
	render IRender
}

// IRepository Interface for basic parsing
type IRepository interface {
	ScrapePrice(url model.Url, selectors model.ParseSelectors, formatter func(model.PriceRecord) model.PriceRecord) (model.PriceRecord, error)
}

type IRender interface {
	Render(html string) ([]byte, error)
}

func New(repo IRepository, render IRender) *Service {
	return &Service{repo, render}
}
