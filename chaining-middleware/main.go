package main

import (
	"net/http"
	// "fmt"
	// "encoding/json"
)

type city struct {
	Name string
	Area uint64
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method == "POST" {
		// tell everything is fine
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - Created"))
	} else {
		// Say method not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not Allowed"))
	}
}

func main() {
	http.HandleFunc("/city", mainHandler)
	http.ListenAndServe(":8080", nil)
}