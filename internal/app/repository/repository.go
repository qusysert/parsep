package repository

import "github.com/gocolly/colly"

type Repository struct {
	colly colly.Collector
}

func New(c colly.Collector) *Repository {
	return &Repository{colly: c}
}
