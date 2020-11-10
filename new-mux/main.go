package main

import (
	"net/http"
	"fmt"
	"math/rand"
)

func main() {
	newMux := http.NewServeMux()
	newMux.HandleFunc("/randomFloat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Random Float: %f\n", rand.Float64())
	})
	newMux.HandleFunc("/randomInt", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Random Int: %v\n", rand.Intn(100))
	})
	http.ListenAndServe(":8080", newMux)
}