package requests

import (
	"fmt"
	"net/http"
)

func HelloHandler(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/hello" {
		http.Error(writer, "404 not found", http.StatusNotFound)
		return
	}
	if req.Method != "GET" {
		http.Error(writer, "Method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "Hello!")
}

func FormHandler(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/form" {
		http.Error(writer, "404 not found", http.StatusNotFound)
	}
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(writer, "req.ParseForm() got err! %v", err)
		return
	}

	fmt.Fprintf(writer, "POST method successful\n\n\n")
	username := req.FormValue("username")
	password := req.FormValue("password")

	if username == "" {
		http.Error(writer, "missing username [username not found]", http.StatusBadRequest)
		return
	}

	if password == "" {
		http.Error(writer, "missing password [password not found]", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(writer, "username: %s\n", username)
	fmt.Fprintf(writer, "-------------------\n")
	fmt.Fprintf(writer, "password: %s\n\n", password)

	if req.Method != "POST" {
		http.Error(writer, "Method not supported", http.StatusNotFound)
	}
	fmt.Fprintf(writer, "Welcome username: %s,your password is %s", username, password)
}
