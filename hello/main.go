package main

import (
	"io"
	"net/http"
	"log"
)

func MyServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!\n")
}

func main() {
	http.HandleFunc("/", MyServer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
