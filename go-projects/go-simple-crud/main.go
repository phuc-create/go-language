package main

import (
	"fmt"
	"log"
	"main/go-projects/go-simple-crud/apis"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	GET    = "GET"
	POST   = "POST"
	PUT    = "PUT"
	DELETE = "DELETE"
)

type ObjSimple interface {
	Checker() bool
}
type Bol struct {
	B bool
}

func (b Bol) Checker() bool {
	return false
}

func main() {

	// var idea ObjSimple
	// if idea != nil {
	// 	fmt.Println("not null")
	// } else {
	// 	fmt.Println(idea)
	// }

	router := mux.NewRouter()
	router.HandleFunc("/api/user/find", apis.FindUser).Methods(GET)
	router.HandleFunc("/api/user/users", apis.GetUsers).Methods(GET)

	// err := http.ListenAndServe(":5050", router)
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("Start server at port 5050\n")
	log.Fatal(http.ListenAndServe(":5050", router))
}
