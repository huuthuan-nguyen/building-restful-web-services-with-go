package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"os"
	"log"
	"net/http"
)

func mainLogic(w http.ResponseWriter, r *http.Request) {
	log.Println("Proccessing request!")
	w.Write([]byte("OK"))
	log.Println("Finished processing request!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", mainLogic)
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	http.ListenAndServe(":8080", loggedRouter)
}