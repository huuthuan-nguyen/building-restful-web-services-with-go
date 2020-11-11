package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

// Article Handler
func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	// mux.Vars returns all path parameters as a map
	vars := mux.Vars(r) // parameters map

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func main() {
	// create a new router
	r := mux.NewRouter()

	// attach a path with handler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)
	server := &http.Server{
		Handler: r,
		Addr: ":8080",
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}
	log.Fatal(server.ListenAndServe())
}