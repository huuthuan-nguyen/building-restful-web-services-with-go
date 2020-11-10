package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

// struct for custom multiplexer
type CustomServeMux struct {

}

// override the handle
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Your random number is: %f", rand.Float64())
}

func main() {
	// struct that implemented ServeHTTP is multiplexer
	mux := &CustomServeMux{}
	http.ListenAndServe(":8080", mux)
}