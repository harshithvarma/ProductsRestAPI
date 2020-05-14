// Package classification of Product API
//
// Documentation for Product API
//
//  Schemes:http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	-application/json
//
//	Produces:
//	-application/json
//swagger:meta
package handlers

import (
	"ProductRestAPI/data"
	"context"
	"fmt"
	"log"
	"net/http"
)

// A list of products returns in the response
//swagger:response productResponse
type ProductsResponseWrapper struct{
	// All products in the system
	//in: body
	body []data.Product
}

// ProductHandles is a http.handler
type ProductHandler struct {
	l *log.Logger
}

//NewProductHandler creates a products handler with given logger
func NewProductHandler(l *log.Logger) *ProductHandler {
	return &ProductHandler{l}
}

// keyProduct is for the middleware to extract the value of the context
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

		// validate the product received from request
		err=prod.Validate()
		if err!=nil{
			p.l.Println("[ERROR] Validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s",err), http.StatusBadRequest)
			return
		}

		//the json object after validation is inserted into the context for further propagation
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)
		next.ServeHTTP(rw, r)

	})
}
