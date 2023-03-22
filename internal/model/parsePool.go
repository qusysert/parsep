package model

var kitcoSelectors = []ParseSelectors{
	{},
}

var mcxSelectors = []ParseSelectors{
	{
		PriceSelector:  ".commodityPrice",
		ChangeSelector: ".perChng",
	},
}

var lmeSelectors = []ParseSelectors{
	{
		PriceSelector:  ".hero-metal-data__number",
		ChangeSelector: ".hero-metal-data__change",
	},
}

var ColorMetalTable = TableDataLinks{
	Title: "Цветные металлы: ",
	ParsePool: []ParsePoint{
		{
			Exchange:  "MCX",
			PriceType: "Price",
			Unit:      "Per 1 KGS",
			Urls: []Url{
				{"Никель", "https://economictimes.indiatimes.com/commoditysummary/symbol-NICKEL.cms"},
				{"Алюминий", "https://economictimes.indiatimes.com/commoditysummary/symbol-ALUMINIUM.cms"},
				{"Цинк", "https://economictimes.indiatimes.com/commoditysummary/symbol-ZINC.cms"},
				{"Олово*", "https://economictimes.indiatimes.com/commoditysummary/symbol-LEAD.cms"},
				{"Медь", "https://economictimes.indiatimes.com/commoditysummary/symbol-COPPER.cms"},
			},
			Selectors: mcxSelectors,
			Formatter: mcxFormatter,
		},
		{
			Exchange:  "LME",
			PriceType: "3-month",
			Unit:      "$",
			Urls: []Url{
				{"Медь", "https://www.lme.com/Metals/Non-ferrous/LME-Copper#Trading+day+summary"},
				{"Алюминий", "https://www.lme.com/en/Metals/Non-ferrous/LME-Aluminium#Trading+day+summary"},
				{"Никель", "https://www.lme.com/en/Metals/Non-ferrous/LME-Nickel"},
				{"Цинк", "https://www.lme.com/en/Metals/Non-ferrous/LME-Zinc#Trading+day+summary"},
				{"Олово*", "https://www.lme.com/Metals/Non-ferrous/LME-Lead#Trading+day+summary"},
			},
			Selectors: lmeSelectors,
			Formatter: lmeFormatter,
		},
	},
}
