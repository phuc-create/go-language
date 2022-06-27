package main

import (
	"fmt"
	"log"
	"main/go-projects/go-server/requests"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", requests.FormHandler)
	http.HandleFunc("/hello", requests.HelloHandler)
	fmt.Printf("starting server at 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
