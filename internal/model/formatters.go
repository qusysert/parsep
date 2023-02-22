package model

import "strings"

func lmeFormatter(pr PriceRecord) PriceRecord {
	m := strings.TrimSpace(strings.ReplaceAll(pr.Material, "\n", ""))
	m = strings.TrimPrefix(m, "LME ")
	pr.Material = m
	return pr
}

func mcxFormatter(pr PriceRecord) PriceRecord {
	m := strings.TrimSpace(strings.ReplaceAll(pr.Material, "\n", ""))
	m = strings.ReplaceAll(m, "&nbsp;", " ")
	m = strings.TrimSuffix(m, "Rate")
	m = m[:len(m)-2]
	pr.Material = m
	return pr
}
