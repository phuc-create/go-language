package routes

import (
	"main/go-projects/go-books-store-BE/pkg/controllers"

	"github.com/gorilla/mux"
	// "github.com/phuc-create/go-language/go-projects/go-books-store-BE/pkg/controllers"
)

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

var RegisterBooksRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetAllBooks).Methods(POST)
}
