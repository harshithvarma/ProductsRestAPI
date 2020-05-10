package main

import (
	"ProductRestAPI/handlers"
	"context"
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"os/signal"
	"time"
	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main(){

	env.Parse()
	l:=log.New(os.Stdout,"product-api",log.LstdFlags)
	s:=mux.NewRouter()
	ph:=handlers.NewProductHandler(l)
	getRouter:=s.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/",ph.GetProducts)

	putRouter:=s.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/{id:[0-9]+}",ph.UpdateProduct)
	putRouter.Use(ph.MiddlewareValidateProduct)
	postRouter:=s.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/",ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)


	server:=http.Server{
		Addr:              *bindAddress,
		Handler:           s,
		TLSConfig:         nil,
		ReadTimeout:       5*time.Second,
		WriteTimeout:      10*time.Second,
		IdleTimeout:       120*time.Second,
	}

	go func(){


		l.Println("Starting the server on port 9090")

		err:=server.ListenAndServe()
		if err!=nil{
			l.Printf("Error starting the server: %s\n",err)
			os.Exit(1)
		}
	}()

	c:=make(chan os.Signal,1)
	signal.Notify(c,os.Interrupt)
	signal.Notify(c,os.Kill)

	sig:=<-c
	log.Println("got signal:",sig)

	ctx,_:=context.WithTimeout(context.Background(),30*time.Second)
	server.Shutdown(ctx)

}
