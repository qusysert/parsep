package model

type PriceRecord struct {
	Material string
	Price    float64
}

type ParseSelectors struct {
	TitleSelector string
	PriceSelector string
}

type ParsePool struct {
	Url       string
	Selectors ParseSelectors
	Formatter func(PriceRecord) PriceRecord
}
