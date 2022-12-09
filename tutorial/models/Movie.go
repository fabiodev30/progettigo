package models

type Movie struct {
	ID        string    `json:"id"`
	Codicefis string    `json:"cdp"`
	Title     string    `json:"titolo"`
	Director  *Director `json:"direttore"`
}
