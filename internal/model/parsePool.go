package model

var PreciousMetalTable = TableDataLinks{
	Title: "Драгоценные металлы: ",
	ParsePool: []ParsePoint{
		{
			Exchange:  "Kitco",
			PriceType: "Price",
			Unit:      "$/oz",
			Urls: []Url{
				{"Золото", "https://www.kitco.com/market/", kitcoGoldSelectors},
				{"Серебро", "https://www.kitco.com/market/", kitcoSilverSelectors},
				{"Платина", "https://www.kitco.com/market/", kitcoPlatinumSelectors},
				{"Палладий", "https://www.kitco.com/market/", kitcoPalladiumSelectors},
			},
			Formatter: kitcoFormatter,
		},
		{
			Exchange:  "Moneymetals",
			PriceType: "Price",
			Unit:      "$/oz",
			Urls: []Url{
				{"Золото", "https://www.moneymetals.com/precious-metals-charts", moneymetalGoldSelectors},
				{"Серебро", "https://www.moneymetals.com/precious-metals-charts", moneymetalSilverSelectors},
				{"Платина", "https://www.moneymetals.com/precious-metals-charts", moneymetalPlatinumSelectors},
				{"Палладий", "https://www.moneymetals.com/precious-metals-charts", moneymetalPalladiumSelectors},
			},
			Formatter: moneymetalFormatter,
		},
	},
}

var ColorMetalTable = TableDataLinks{
	Title: "Цветные металлы: ",
	ParsePool: []ParsePoint{
		{
			Exchange:  "MCX",
			PriceType: "Price",
			Unit:      "₹/KGS",
			Urls: []Url{
				{"Алюминий", "https://economictimes.indiatimes.com/commoditysummary/symbol-ALUMINIUM.cms", mcxSelectors},
				{"Медь", "https://economictimes.indiatimes.com/commoditysummary/symbol-COPPER.cms", mcxSelectors},
				{"Цинк", "https://economictimes.indiatimes.com/commoditysummary/symbol-ZINC.cms", mcxSelectors},
				{"Никель", "https://economictimes.indiatimes.com/commoditysummary/symbol-NICKEL.cms", mcxSelectors},
				{"Олово*", "https://economictimes.indiatimes.com/commoditysummary/symbol-LEAD.cms", mcxSelectors},
			},
			Formatter: mcxFormatter,
		},
		{
			Exchange:  "LME",
			PriceType: "3-month",
			Unit:      "$",
			Urls: []Url{
				{"Алюминий", "https://www.lme.com/en/Metals/Non-ferrous/LME-Aluminium#Trading+day+summary", lmeSelectors},
				{"Медь", "https://www.lme.com/Metals/Non-ferrous/LME-Copper#Trading+day+summary", lmeSelectors},
				{"Цинк", "https://www.lme.com/en/Metals/Non-ferrous/LME-Zinc#Trading+day+summary", lmeSelectors},
				{"Никель", "https://www.lme.com/en/Metals/Non-ferrous/LME-Nickel", lmeSelectors},
				{"Олово*", "https://www.lme.com/Metals/Non-ferrous/LME-Tin#Trading+day+summary", lmeSelectors},
			},
			Formatter: lmeFormatter,
		},
	},
}
