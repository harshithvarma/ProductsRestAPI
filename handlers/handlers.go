package handlers

import (
	"ProductRestAPI/data"
	"context"
	"fmt"
	"log"
	"net/http"
)

type ProductHandler struct {
	l *log.Logger
}

func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

type KeyProduct struct{}

func (p *ProductHandler) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		err := prod.FromJson(r.Body)
		if err != nil {
			p.l.Println("[ERROR] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		err=prod.Validate()
		if err!=nil{
			p.l.Println("[ERROR] Validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s",err), http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)

	})
}
