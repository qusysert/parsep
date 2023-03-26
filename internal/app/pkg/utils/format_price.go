package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func FormatPrice(p string) (float64, error) {
	if strings.Contains(p, ",") && strings.Contains(p, ".") {
		p = strings.ReplaceAll(p, ",", "")
	} else if strings.Contains(p, ",") {
		p = strings.ReplaceAll(p, ",", ".")
	}
	price, err := strconv.ParseFloat(regexp.MustCompile(`[^\d\.]`).ReplaceAllString(p, ""), 64)
	return price, err
}
