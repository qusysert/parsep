package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator"
	"io/ioutil"
	"net/http"
	"parser/internal/model"
	"reflect"
)

type IService interface {
	FormTableContent(tableData model.TableDataLinks) (model.TabledData, error)
	GenTableHTML(data model.TabledData) (string, error)
	RenderHtml(html string) ([]byte, error)
}

type Handler struct {
	service IService
}

func New(srv IService) *Handler {
	return &Handler{service: srv}
}

// Unmarshal request, do work fn(), then marshall response into JSON anf return
func handle[REQ any, RESP any](w http.ResponseWriter, r *http.Request, fn func(req REQ) (RESP, error)) {
	var req REQ
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusBadRequest)
		return
	}

	if isNil(req) != true {
		err = json.Unmarshal(body, &req)
		if err != nil {
			http.Error(w, fmt.Sprintf("%+v", err), http.StatusBadRequest)
			return
		}
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusBadRequest)
		return
	}

	resp, err := fn(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusInternalServerError)
		return
	}

	respJSON, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusInternalServerError)
		return
	}

	_, err = w.Write(respJSON)
	if err != nil {
		http.Error(w, fmt.Sprintf("%+v", err), http.StatusInternalServerError)
		return
	}
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Array, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
