package model

var moneymetalGoldSelectors = []ParseSelectors{
	{
		PriceSelector:  "tbody tr:nth-child(1) td:nth-child(2)",
		ChangeSelector: "tbody tr:nth-child(1) td:nth-child(4) div",
	},
}
var moneymetalSilverSelectors = []ParseSelectors{
	{
		PriceSelector:  "tbody tr:nth-child(2) td:nth-child(2)",
		ChangeSelector: "tbody tr:nth-child(2) td:nth-child(4) div",
	},
}

var moneymetalPlatinumSelectors = []ParseSelectors{
	{
		PriceSelector:  "tbody tr:nth-child(3) td:nth-child(2)",
		ChangeSelector: "tbody tr:nth-child(3) td:nth-child(4) div",
	},
}

var moneymetalPalladiumSelectors = []ParseSelectors{
	{
		PriceSelector:  "tbody tr:nth-child(4) td:nth-child(2)",
		ChangeSelector: "tbody tr:nth-child(4) td:nth-child(4) div",
	},
}
