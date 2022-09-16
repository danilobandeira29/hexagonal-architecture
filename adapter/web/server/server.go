package server

import (
	"github.com/codegangsta/negroni"
	"github.com/danilobandeira29/hexagonal-architecture/adapter/web/handler"
	"github.com/danilobandeira29/hexagonal-architecture/application"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type Webserver struct {
	service application.ProductServiceInterface
}

func NewWebserver(service application.ProductServiceInterface) *Webserver {
	return &Webserver{
		service,
	}
}

func (w *Webserver) Serve() {
	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())
	handler.MakeProductHandlers(r, n, w.service)
	http.Handle("/", r)
	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
