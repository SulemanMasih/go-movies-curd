package main

import (
	"encoding/json" // Package for JSON encoding and decoding
	"fmt"           // Package for formatted I/O
	"log"           // Package for logging
	"math/rand"     // Package for generating random numbers
	"net/http"      // Package for HTTP client and server implementations
	"strconv"       // Package for string conversions
	"github.com/gorilla/mux" // Package for creating HTTP routers
)

// Movie struct represents a movie with an ID, ISBN, Title, and Director
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

// Director struct represents a director with a first and last name
type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// movies slice to store movie data
var movies []Movie

// main function initializes the movies and starts the HTTP server
func main() {
	fmt.Println("hello")

	// Adding initial movie data
	movies = append(movies, Movie{ID: "1", Isbn: "45464", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Smith"}})
	movies = append(movies, Movie{ID: "2", Isbn: "43364", Title: "Movie Two", Director: &Director{Firstname: "Wood", Lastname: "Joe"}})

	// Creating a new router
	r := mux.NewRouter()

	// Route handlers / endpoints
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	// Starting the server on port 8000
	fmt.Printf("Starting the server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// getMovies handles the GET request to retrieve all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// deleteMovie handles the DELETE request to delete a movie by ID
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get route parameters
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Remove the movie from the slice
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

// getMovie handles the GET request to retrieve a movie by ID
func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get route parameters
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createMovie handles the POST request to create a new movie
func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the request body into movie struct
	movie.ID = strconv.Itoa(rand.Intn(1000))   // Generate a random ID for the new movie
	movies = append(movies, movie)             // Add the new movie to the slice
	json.NewEncoder(w).Encode(movie)           // Encode the new movie as JSON and send the response
}

// updateMovie handles the PUT request to update a movie by ID
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get route parameters
	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...) // Remove the old movie from the slice
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie) // Decode the request body into movie struct
			movie.ID = params["id"]                    // Assign the same ID to the updated movie
			movies = append(movies, movie)             // Add the updated movie to the slice
			json.NewEncoder(w).Encode(movie)           // Encode the updated movie as JSON and send the response
			return
		}
	}
}
