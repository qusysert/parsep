package handler

import (
	"fmt"
	"net/http"
	"parser/internal/model"
)

type GetTableRequest struct {
	TableName string `json:"table_name"`
}
type GetTableResponse struct {
	Success bool `json:"success"`
}

func (h Handler) GetTableHandler(w http.ResponseWriter, r *http.Request) {
	handle(w, r, func(req GetTableRequest) (GetTableResponse, error) {
		var content model.TabledData
		var err error
		switch req.TableName {
		case "ColorMetalTable":
			content, err = h.service.FormTableContent(model.ColorMetalTable)
			if err != nil {
				return GetTableResponse{Success: false}, err
			}
		default:
			return GetTableResponse{Success: false}, fmt.Errorf("table name %s not found", req.TableName)
		}

		htmlString, err := h.service.GenTableHTML(content)
		if err != nil {
			return GetTableResponse{Success: false}, err
		}
		imageBytes, err := h.service.RenderHtml(htmlString)
		if err != nil {
			return GetTableResponse{Success: false}, err
		}
		w.Header().Set("Content-Type", "image/png")

		_, _ = w.Write(imageBytes)
		return GetTableResponse{Success: true}, nil
	})

}
