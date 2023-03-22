package model

import (
	"regexp"
	"strings"
)

func lmeFormatter(pr PriceRecord) PriceRecord {
	pr.Change = regexp.MustCompile(`[^\d\.\+\-]`).ReplaceAllString(pr.Change, "")
	if pr.Change[0:1] != "+" && pr.Change[0:1] != "-" {
		pr.Change = "+" + pr.Change
	}
	return pr
}

func mcxFormatter(pr PriceRecord) PriceRecord {
	// abs (percent)
	cArr := strings.Split(pr.Change, " ")
	pr.Change = regexp.MustCompile(`[^\d\.\+\-]`).ReplaceAllString(cArr[1], "")
	pr.Change = regexp.MustCompile(`[^\d\.\+\-]`).ReplaceAllString(pr.Change, "")
	if pr.Change[0:1] != "+" && pr.Change[0:1] != "-" {
		pr.Change = "+" + pr.Change
	}
	return pr
}
