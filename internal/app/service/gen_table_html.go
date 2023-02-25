package service

import (
	"fmt"
	"parser/internal/model"
	"strings"
)

func (s *Service) GenTableHTML(data []model.TableColumn) string {
	var sb strings.Builder

	// start table
	sb.WriteString("<table>")

	// write headers
	sb.WriteString("<tr><th></th>")
	for _, metal := range data {
		fmt.Fprintf(&sb, "<th>%s</th>", metal.Title)
	}
	sb.WriteString("</tr>")

	// write data rows
	for _, material := range data[0].Prices {
		fmt.Fprintf(&sb, "<tr><td>%s</td>", material.Material)
		for _, metal := range data {
			var price float64
			for _, m := range metal.Prices {
				if m.Material == material.Material {
					price = m.Price
					break
				}
			}
			fmt.Fprintf(&sb, "<td>%.2f</td>", price)
		}
		sb.WriteString("</tr>")
	}

	// end table
	sb.WriteString("</table>")

	return sb.String()
}
