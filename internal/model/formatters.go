package model

import (
	"regexp"
	"strings"
)

func lmeFormatter(pr PriceRecord) PriceRecord {
	m := strings.TrimSpace(strings.ReplaceAll(pr.Material, "\n", ""))
	m = strings.TrimPrefix(m, "LME ")
	pr.Material = m
	pr.Change = regexp.MustCompile(`[^\d\.\+\-]`).ReplaceAllString(pr.Change, "")
	return pr
}

func mcxFormatter(pr PriceRecord) PriceRecord {
	m := strings.TrimSpace(strings.ReplaceAll(pr.Material, "\n", ""))
	m = strings.ReplaceAll(m, "&nbsp;", " ")
	m = strings.TrimSuffix(m, "Rate")
	m = m[:len(m)-2]
	pr.Material = m
	// abs (percent)
	cArr := strings.Split(pr.Change, " ")
	pr.Change = regexp.MustCompile(`[^\d\.\+\-]`).ReplaceAllString(cArr[1], "")
	return pr
}
