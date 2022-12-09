package operazioni

import (
	"fmt"
	"net/http"
	"strings"
)

var Somma = func(w http.ResponseWriter, r *http.Request) {
	stringNumeri := r.URL.Query()["n"]
	percorso := strings.Split(r.URL.Path, "/")
	var risultato int
	if len(stringNumeri) > 0 {
		risultato = Operazione(percorso[1], stringNumeri)
	}
	risposta := fmt.Sprintf("La %s dei numeri è:%d", percorso[1], risultato)
	w.Write([]byte(risposta))
}
