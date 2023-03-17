package model

type PriceRecord struct {
	Material string
	Price    float64
	Change   string
}

type ParseSelectors struct {
	TitleSelector  string
	PriceSelector  string
	ChangeSelector string
}

type ParsePoint struct {
	Exchange  string
	PriceType string
	Unit      string
	Urls      []string
	Selectors ParseSelectors
	Formatter func(PriceRecord) PriceRecord
}

type TableColumn struct {
	Title  string
	Prices []PriceRecord
}
