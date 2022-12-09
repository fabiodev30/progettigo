package operazioni

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var Differenza = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sottraendo, err := strconv.Atoi(vars["sottraendo"])
	if err != nil {
		return
	}
	sottrattore, err := strconv.Atoi(vars["sottrattore"])
	if err != nil {
		return
	}
	differenza := sottraendo - sottrattore
	risposta := fmt.Sprintf("La differenza fra i numeri è:%d", differenza)
	w.Write([]byte(risposta))
}
