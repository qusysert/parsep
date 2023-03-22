package model

var kitcoGoldSelectors = []ParseSelectors{
	{
		PriceSelector:  ".AU-ask",
		ChangeSelector: ".AU-change-percent .spotred",
	},
}

var kitcoSilverSelectors = []ParseSelectors{
	{
		PriceSelector:  ".AG-ask",
		ChangeSelector: ".AG-change-percent .spotred",
	},
}

var kitcoPlatinumSelectors = []ParseSelectors{
	{
		PriceSelector:  ".PT-ask",
		ChangeSelector: ".PT-change-percent .spotred",
	},
}

var kitcoPalladiumSelectors = []ParseSelectors{
	{
		PriceSelector:  ".PD-ask",
		ChangeSelector: ".PD-change-percent .spotred",
	},
}

var kitcoRhodiumSelectors = []ParseSelectors{
	{
		PriceSelector:  ".RH-ask",
		ChangeSelector: ".RH-change-percent .spotred",
	},
}
