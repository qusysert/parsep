package model

var cbrGoldSelectors = ParseSelectors{
	PriceSelector:     "table.data > tbody > tr:nth-of-type(2) > td:nth-of-type(2)",
	PrevPriceSelector: "table.data > tbody > tr:nth-of-type(3) > td:nth-of-type(2)",
}

var cbrSilverSelectors = ParseSelectors{
	PriceSelector:     "table.data > tbody > tr:nth-of-type(2) > td:nth-of-type(3)",
	PrevPriceSelector: "table.data > tbody > tr:nth-of-type(3) > td:nth-of-type(3)",
}

var cbrPlatinumSelectors = ParseSelectors{
	PriceSelector:     "table.data > tbody > tr:nth-of-type(2) > td:nth-of-type(4)",
	PrevPriceSelector: "table.data > tbody > tr:nth-of-type(3) > td:nth-of-type(4)",
}

var cbrPalladiumSelectors = ParseSelectors{
	PriceSelector:     "table.data > tbody > tr:nth-of-type(2) > td:nth-of-type(5)",
	PrevPriceSelector: "table.data > tbody > tr:nth-of-type(3) > td:nth-of-type(5)",
}
