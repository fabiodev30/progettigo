package main

import (
	"bookstore/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	routes.NegozioLibriRotte(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
