package model

var PreciousMetalTable = TableDataLinks{
	Title: "Драгоценные металлы: ",
	ParsePool: []ParsePoint{
		{
			Exchange:  "Kitco",
			PriceType: "Price",
			Unit:      "$/кг",
			Urls: []Url{
				{"Золото", "https://www.kitco.com/market/", kitcoGoldSelectors},
				{"Серебро", "https://www.kitco.com/market/", kitcoSilverSelectors},
				{"Платина", "https://www.kitco.com/market/", kitcoPlatinumSelectors},
				{"Палладий", "https://www.kitco.com/market/", kitcoPalladiumSelectors},
				{"Родий", "https://www.kitco.com/market/", kitcoRhodiumSelectors},
			},
			Formatter: kitcoFormatter,
		},
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
				{"Никель", "https://economictimes.indiatimes.com/commoditysummary/symbol-NICKEL.cms", mcxSelectors},
				{"Алюминий", "https://economictimes.indiatimes.com/commoditysummary/symbol-ALUMINIUM.cms", mcxSelectors},
				{"Цинк", "https://economictimes.indiatimes.com/commoditysummary/symbol-ZINC.cms", mcxSelectors},
				{"Олово*", "https://economictimes.indiatimes.com/commoditysummary/symbol-LEAD.cms", mcxSelectors},
				{"Медь", "https://economictimes.indiatimes.com/commoditysummary/symbol-COPPER.cms", mcxSelectors},
			},
			Formatter: mcxFormatter,
		},
		{
			Exchange:  "LME",
			PriceType: "3-month",
			Unit:      "$",
			Urls: []Url{
				{"Медь", "https://www.lme.com/Metals/Non-ferrous/LME-Copper#Trading+day+summary", lmeSelectors},
				{"Алюминий", "https://www.lme.com/en/Metals/Non-ferrous/LME-Aluminium#Trading+day+summary", lmeSelectors},
				{"Никель", "https://www.lme.com/en/Metals/Non-ferrous/LME-Nickel", lmeSelectors},
				{"Цинк", "https://www.lme.com/en/Metals/Non-ferrous/LME-Zinc#Trading+day+summary", lmeSelectors},
				{"Олово*", "https://www.lme.com/Metals/Non-ferrous/LME-Lead#Trading+day+summary", lmeSelectors},
			},
			Formatter: lmeFormatter,
		},
	},
}
