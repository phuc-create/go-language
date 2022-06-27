package main

import (
	"fmt"
	"log"
	moviescontroller "main/go-projects/go-movies-crud/moviesController"
	"net/http"

	"github.com/gorilla/mux"
)

// type Movies struct {
// 	ID       string    `json:"id"`
// 	Title    string    `json:"title"`
// 	Director *Director `json:"director"`
// }
//
// type Director struct {
// 	ID   string `json:"id"`
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }
//
// var MoviesList []Movies

func main() {
	router := mux.NewRouter()
	moviescontroller.MoviesList = append(moviescontroller.MoviesList, moviescontroller.Movies{ID: "1", Title: "KingKong", Director: &moviescontroller.Director{ID: "1", Name: "Sam", Age: 22}})
	moviescontroller.MoviesList = append(moviescontroller.MoviesList, moviescontroller.Movies{ID: "2", Title: "Kong Island", Director: &moviescontroller.Director{ID: "1", Name: "Sam", Age: 22}})
	router.HandleFunc("/movies", moviescontroller.GetAllMovies).Methods("GET")
	router.HandleFunc("/movies/{id}", moviescontroller.GetMoviesById).Methods("GET")
	router.HandleFunc("/movies", moviescontroller.CreateMovie).Methods("POST")
	router.HandleFunc("/movies/{id}", moviescontroller.UpdateMoviesById).Methods("PUT")
	router.HandleFunc("/movies/{id}", moviescontroller.DeleteMoviesById).Methods("DELETE")

	fmt.Printf("Start server at port 7070\n")

	log.Fatal(http.ListenAndServe(":7070", router))
}
