package operazioni

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var Moltiplicazione = func(w http.ResponseWriter, r *http.Request) {
	stringNumeri := r.URL.Query()["n"]
	percorso := strings.Split(r.URL.Path, "/")
	if len(stringNumeri) > 0 {
		// faccio questa cosa di prendere il primo numero e poi saltarlo nel for dato che se prendevo una variabile esterna me l ainizializzava a 0 e quindi mi restituiva risultato 0
		moltiplicazione := 1
		for i := range stringNumeri {
			num, err := strconv.Atoi(stringNumeri[i])
			if err != nil {
				return
			}
			moltiplicazione *= num
		}
		risposta := fmt.Sprintf("La %s dei numeri è:%d", percorso[1], moltiplicazione)
		w.Write([]byte(risposta))
	}

}
