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
			},
			Formatter: kitcoFormatter,
		},
		{
			Exchange:  "Moneymetals",
			PriceType: "Price",
			Unit:      "$/кг",
			Urls: []Url{
				{"Золото", "https://www.moneymetals.com/precious-metals-charts", moneymetalGoldSelectors},
				{"Серебро", "https://www.moneymetals.com/precious-metals-charts", moneymetalSilverSelectors},
				{"Платина", "https://www.moneymetals.com/precious-metals-charts", moneymetalPlatinumSelectors},
				{"Палладий", "https://www.moneymetals.com/precious-metals-charts", moneymetalPalladiumSelectors},
			},
			Formatter: moneymetalFormatter,
		},
		{
			Exchange:  "ЦБ",
			PriceType: "Price",
			Unit:      "₽/г",
			Urls: []Url{
				{"Золото", "https://www.cbr.ru/hd_base/metall/metall_base_new/", cbrGoldSelectors},
				{"Серебро", "https://www.cbr.ru/hd_base/metall/metall_base_new/", cbrSilverSelectors},
				{"Платина", "https://www.cbr.ru/hd_base/metall/metall_base_new/", cbrPlatinumSelectors},
				{"Палладий", "https://www.cbr.ru/hd_base/metall/metall_base_new/", cbrPalladiumSelectors},
			},
			Formatter: cbrFormatter,
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
				{"Свинец", "https://economictimes.indiatimes.com/commoditysummary/symbol-LEAD.cms", mcxSelectors},
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
				{"Свинец", "https://www.lme.com/Metals/Non-ferrous/LME-Lead#Trading+day+summary", lmeSelectors},
			},
			Formatter: lmeFormatter,
		},
	},
}
