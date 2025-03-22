package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json="id"`
	ISBN     string    `json="isbn"`
	Title    string    `json="title"`
	Director *Director `json="director"`
}

type Director struct {
	Firstname string `json="firstname"`
	Lastname  string `json="lastname"`
}

var movies []Movie

func main() {
	initialData()

	router := mux.NewRouter()
	// routes
	router.HandleFunc("/movies", getMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	router.HandleFunc("/movies", createMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	router.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// deploying local server
	http.ListenAndServe(":8080", router)
}

func initialData() {
	movies = append(movies, Movie{ID: "1", ISBN: "0101", Title: "Movie 1", Director: &Director{Firstname: "Name 1", Lastname: "Lastname 1"}})
	movies = append(movies, Movie{ID: "2", ISBN: "0202", Title: "Movie 2", Director: &Director{Firstname: "Name 2", Lastname: "Lastname 2"}})
	movies = append(movies, Movie{ID: "3", ISBN: "0303", Title: "Movie 3", Director: &Director{Firstname: "Name 3", Lastname: "Lastname 3"}})
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, movie := range movies {
		if movie.ID == params["id"] {
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	response := map[string]string{"message": "Movie not found"}
	json.NewEncoder(w).Encode(response)
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	movies = append(movies, newMovie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)

	for i, movie := range movies {
		if movie.ID == params["id"] {
			movies[i] = updatedMovie
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for i, movie := range movies {
		if movie.ID == params["id"] {
			movies = append(movies[:i], movies[i+1:]...)
			break
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
