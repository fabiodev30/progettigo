package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"tutorial/models"

	"github.com/gorilla/mux"
)

var Movies []models.Movie

func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for index, v := range Movies {
		if v.ID == id {
			Movies = append(Movies[:index], Movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(Movies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(movie)

}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var movie models.Movie
	json.NewDecoder(r.Body).Decode(&movie)
	for _, v := range Movies {
		if v.ID == id {
			movie = v
			// json.NewEncoder(w).Encode(movie)
			// break
		}
	}
	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for index, v := range Movies {
		if v.ID == id {
			Movies = append(Movies[:index], Movies[index+1:]...)
			var movie models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = id
			Movies = append(Movies, movie)
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}

func main() {
	r := mux.NewRouter()
	Movies = append(Movies, models.Movie{ID: "1", Codicefis: "qsgqgsjuq", Title: "Spiderman", Director: &models.Director{Nome: "Pippo", Cognome: "Baudo"}})
	Movies = append(Movies, models.Movie{ID: "2", Codicefis: "wdghw", Title: "Batman", Director: &models.Director{Nome: "Steve", Cognome: "Jobs"}})
	r.HandleFunc("/movies", GetMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")
	r.HandleFunc("/movies", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", DeleteMovie).Methods("DELETE")
	fmt.Println("Startando server at port 8000")
	http.ListenAndServe(":8000", r)
}
