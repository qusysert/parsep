package model

var kitcoGoldSelectors = []ParseSelectors{
	{
		PriceSelector:  ".odd #AU-ask",
		ChangeSelector: "#AU-change-percent",
	},
}

var kitcoSilverSelectors = []ParseSelectors{
	{
		PriceSelector:  ".even #AG-ask",
		ChangeSelector: "#AG-change-percent",
	},
}

var kitcoPlatinumSelectors = []ParseSelectors{
	{
		PriceSelector:  ".odd #PT-ask",
		ChangeSelector: "#PT-change-percent",
	},
}

var kitcoPalladiumSelectors = []ParseSelectors{
	{
		PriceSelector:  ".even #PD-ask",
		ChangeSelector: "#PD-change-percent",
	},
}

var kitcoRhodiumSelectors = []ParseSelectors{
	{
		PriceSelector:  ".odd #RH-ask",
		ChangeSelector: "#RH-change-percent",
	},
}
