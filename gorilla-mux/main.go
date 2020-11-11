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

func AuthorArticle(w http.ResponseWriter, r *http.Request) {
	// mux.Vars returns all path parameters as a map
	vars := mux.Vars(r) // parameters map

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Author is: %v\n", vars["name"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

// settings handler
func settingsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // parameters map

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Setting ID: %v\n", vars["id"])
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // parameters map

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Detail ID: %v\n", vars["id"])
}

func main() {
	// create a new router
	r := mux.NewRouter()

	// attach a path with handler
	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	// define preceding route
	r.Path("/authors/{name}/{id:[0-9]+}").HandlerFunc(AuthorArticle)

	// subrouter, useful for grouping router
	s := r.PathPrefix("/config").Subrouter()
	s.HandleFunc("/{id:[0-9]+}/settings", settingsHandler)
	s.HandleFunc("/{id:[0-9]+}/details", detailsHandler)

	// serve static with path prefix
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	server := &http.Server{
		Handler: r,
		Addr: ":8080",
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}
	log.Fatal(server.ListenAndServe())
}