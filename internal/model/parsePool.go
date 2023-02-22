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
		Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-NICKEL.cms",
		Selectors: mcxSelectors,
		Formatter: mcxFormatter,
	},
	{
		Exchange:  "MCX",
		Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-ALUMINIUM.cms",
		Selectors: mcxSelectors,
		Formatter: mcxFormatter,
	},
	{
		Exchange:  "MCX",
		Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-ZINC.cms",
		Selectors: mcxSelectors,
		Formatter: mcxFormatter,
	},
	{
		Exchange:  "MCX",
		Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-LEAD.cms",
		Selectors: mcxSelectors,
		Formatter: mcxFormatter,
	},
	{
		Exchange:  "MCX",
		Url:       "https://economictimes.indiatimes.com/commoditysummary/symbol-COPPER.cms",
		Selectors: mcxSelectors,
		Formatter: mcxFormatter,
	},
	{
		Exchange:  "LME",
		Url:       "https://www.lme.com/Metals/Non-ferrous/LME-Copper#Trading+day+summary",
		Selectors: lmeSelectors,
		Formatter: lmeFormatter,
	},
	{
		Exchange:  "LME",
		Url:       "https://www.lme.com/en/Metals/Non-ferrous/LME-Aluminium#Trading+day+summary",
		Selectors: lmeSelectors,
		Formatter: lmeFormatter,
	},
	{
		Exchange:  "LME",
		Url:       "https://www.lme.com/en/Metals/Non-ferrous/LME-Nickel",
		Selectors: lmeSelectors,
		Formatter: lmeFormatter,
	},
	{
		Exchange:  "LME",
		Url:       "https://www.lme.com/en/Metals/Non-ferrous/LME-Zinc#Trading+day+summary",
		Selectors: lmeSelectors,
		Formatter: lmeFormatter,
	},
}
