package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"time"
)

func main(){
	l:=log.New(os.Stdout,"product-api",log.LstdFlags)
	s:=mux.NewRouter()
	getRouter:=s.Methods(http.MethodGet).Subrouter()
	putRouter:=s.Methods(http.MethodPut).Subrouter()
	postRouter:=s.Methods(http.MethodPost).Subrouter()
	delRouter:=s.Methods(http.MethodDelete).Subrouter()

	server:=http.Server{
		Addr:              ":9090",
		Handler:           s,
		TLSConfig:         nil,
		ReadTimeout:       5*time.Second,
		WriteTimeout:      10*time.Second,
		IdleTimeout:       120*time.Second,
	}

}
