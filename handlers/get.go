package handlers

import (
	"ProductRestAPI/data"
	"net/http"
)

func (p *ProductHandler) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")
	list := data.GetProductsList()
	err := list.ToJson(rw)
	if err != nil {
		http.Error(rw, "unable to unmarshal to json", http.StatusInternalServerError)
	}
}
