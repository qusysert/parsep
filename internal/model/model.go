package model

type PriceRecord struct {
	Exchange string
	Material string
	Price    float64
}

type ParseSelectors struct {
	TitleSelector string
	PriceSelector string
}

type ParsePoint struct {
	Exchange  string
	Url       string
	Selectors ParseSelectors
	Formatter func(PriceRecord) PriceRecord
}
