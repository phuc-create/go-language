package moviescontroller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

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
// var movies []Movies

func GetAllMovies(w http.ResponseWriter, req *http.Request) {

	if req.URL.Path != "/movies" {
		http.Error(w, "404 not found the page", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(MoviesList)

}
func GetMoviesById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range MoviesList {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/movies" {
		http.Error(w, "404 not found the page", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var movie Movies
	_ = json.NewDecoder(req.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	MoviesList = append(MoviesList, movie)
	json.NewEncoder(w).Encode(MoviesList)
}
func UpdateMoviesById(w http.ResponseWriter, req *http.Request) {
	// if req.URL.Path != "/movies" {
	// 	http.Error(w, "404 not found the page", http.StatusNotFound)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range MoviesList {
		if item.ID == params["id"] {
			MoviesList = append(MoviesList[:index], MoviesList[index+1:]...)
			var movie Movies
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = strconv.Itoa(rand.Intn(1000000))
			MoviesList = append(MoviesList, movie)
			json.NewEncoder(w).Encode(MoviesList)
		}
	}
}
func DeleteMoviesById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range MoviesList {
		if item.ID == params["id"] {
			MoviesList = append(MoviesList[:index], MoviesList[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(MoviesList)
}
