package operazioni

import "strconv"

// funzione operazione prende come argomenti una stringa con nome operazione e la serie di numeri su cui applicarla e restituisce risultato
func Operazione(nomeOperazione string, numeri []string) (risultato int) {
	for i := range numeri {
		num, err := strconv.Atoi(numeri[i])
		if err != nil {
			return
		}
		if i == 0 {
			risultato = num
			continue
		}
		switch nomeOperazione {
		case "somma":
			risultato += num
		case "moltiplicazione":
			risultato *= num
		}
	}
	return risultato
}
