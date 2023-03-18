package handler

import (
	"net/http"
	"parser/internal/model"
)

func (h Handler) GetTableHandler(w http.ResponseWriter, r *http.Request) {
	content, err := h.service.FormTableContent(model.ParsePool)
	if err != nil {
		http.Error(w, "cant form table content: "+err.Error(), http.StatusInternalServerError)
		return
	}
	htmlString, err := h.service.GenTableHTML(content)
	if err != nil {
		http.Error(w, "cant get table html: "+err.Error(), http.StatusInternalServerError)
		return
	}
	imageBytes, err := h.service.RenderHtml(htmlString)
	if err != nil {
		http.Error(w, "cant render table html to image: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "image/png")

	_, _ = w.Write(imageBytes)
}
