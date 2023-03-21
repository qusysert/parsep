package service

import (
	"fmt"
	"parser/internal/model"
	"strings"
)

func (s *Service) GenTableHTML(data model.TabledData) (string, error) {
	var sb strings.Builder

	// start table
	sb.WriteString("<table>")
	_, err := fmt.Fprintf(&sb, `<caption style="text-align: center; bold">%s</caption>`, data.TableTitle)
	if err != nil {
		return "", fmt.Errorf("cant write table title: %w", err)
	}

	// write headers
	sb.WriteString("<tr><th></th>")
	for _, metal := range data.Columns {
		_, err := fmt.Fprintf(&sb, "<th>%s</th>", metal.Title)
		if err != nil {
			return "", fmt.Errorf("cant format material title: %w", err)
		}
		_, err = fmt.Fprintf(&sb, "<th>%s</th>", "Изм. %")
		if err != nil {
			return "", fmt.Errorf("cant format change title: %w", err)
		}
	}
	sb.WriteString("</tr>")

	// write data rows
	for _, material := range data.Columns[0].Prices {
		_, err := fmt.Fprintf(&sb, "<tr><td>%s</td>", material.Material)
		if err != nil {
			return "", fmt.Errorf("cant format row: %w", err)
		}
		for _, metal := range data.Columns {
			var price float64
			var change string
			for _, m := range metal.Prices {
				if m.Material == material.Material {
					price = m.Price
					change = m.Change
					break
				}
			}
			_, err := fmt.Fprintf(&sb, "<td>%.2f</td>", price)
			if err != nil {
				return "", fmt.Errorf("cant format row price: %w", err)
			}
			_, err = fmt.Fprintf(&sb, "<td>%s</td>", change)
			if err != nil {
				return "", fmt.Errorf("cant format row change: %w", err)
			}
		}
		sb.WriteString("</tr>")
	}

	// end table
	sb.WriteString("</table>")

	return sb.String(), nil
}
