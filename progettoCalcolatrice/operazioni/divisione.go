package operazioni

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Divisione = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dividendo, err := strconv.Atoi(vars["dividendo"])
	if err != nil {
		return
	}
	divisore, err := strconv.Atoi(vars["divisore"])
	if err != nil {
		return
	}
	quoziente := dividendo / divisore
	resto := dividendo % divisore
	risposta := fmt.Sprintf("Il quoziente è %d e il resto è %d", quoziente, resto)
	w.Write([]byte(risposta))
}
