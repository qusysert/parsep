package model

var mcxSelectors = ParseSelectors{
	TitleSelector: ".commodityTitle",
	PriceSelector: ".commodityPrice",
}
var lmeSelectors = ParseSelectors{
	TitleSelector: ".hero-banner__title",
	PriceSelector: ".hero-metal-data__number",
}

var ParsePool = []ParsePoint{
	{
		Exchange:  "MCX",
		PriceType: "Price",
		Unit:      "Per 1 KGS",
		Urls: []string{
			"https://economictimes.indiatimes.com/commoditysummary/symbol-NICKEL.cms",
			"https://economictimes.indiatimes.com/commoditysummary/symbol-ALUMINIUM.cms",
			"https://economictimes.indiatimes.com/commoditysummary/symbol-ZINC.cms",
			"https://economictimes.indiatimes.com/commoditysummary/symbol-LEAD.cms",
			"https://economictimes.indiatimes.com/commoditysummary/symbol-COPPER.cms",
		},
		Selectors: mcxSelectors,
		Formatter: mcxFormatter,
	},
	{
		Exchange:  "LME",
		PriceType: "3-month",
		Unit:      "$",
		Urls: []string{
			"https://www.lme.com/Metals/Non-ferrous/LME-Copper#Trading+day+summary",
			"https://www.lme.com/en/Metals/Non-ferrous/LME-Aluminium#Trading+day+summary",
			"https://www.lme.com/en/Metals/Non-ferrous/LME-Nickel",
			"https://www.lme.com/en/Metals/Non-ferrous/LME-Zinc#Trading+day+summary",
		},
		Selectors: lmeSelectors,
		Formatter: lmeFormatter,
	},
}
