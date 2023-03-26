package model

type PriceRecord struct {
	Material string
	Price    float64
	Change   string
}

type ParseSelectors struct {
	PriceSelector     string
	ChangeSelector    string
	PrevPriceSelector string
}

type ParsePoint struct {
	Exchange  string
	PriceType string
	Unit      string
	Urls      []Url
	Formatter func(PriceRecord) PriceRecord
}

type Url struct {
	Material  string
	Link      string
	Selectors ParseSelectors
}

type TableDataLinks struct {
	Title     string
	ParsePool []ParsePoint
}

type TableColumn struct {
	Title  string
	Prices []PriceRecord
}

type TabledData struct {
	TableTitle string
	Columns    []TableColumn
}
