package main

import (
	"net/http"
	"progettoserver/operazioni"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.Handle("/somma", http.HandlerFunc(operazioni.Somma)).Methods("POST")
	router.Handle("/moltiplicazione", http.HandlerFunc(operazioni.Somma)).Methods("POST")
	router.Handle("/differenza/{sottraendo}/{sottrattore}", http.HandlerFunc(operazioni.Differenza)).Methods("POST")
	router.Handle("/divisione/{dividendo}/{divisore}", http.HandlerFunc(operazioni.Divisione)).Methods("POST")
	// in questo caso tutte le richieste che faremo a localhost verranno passate all'unica funzione che c'è che in questo caso è quella che segue
	err := http.ListenAndServe(":80", router)
	if err != nil {
		return
	}
}
