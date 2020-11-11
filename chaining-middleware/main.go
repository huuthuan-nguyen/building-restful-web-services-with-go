package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

type city struct {
	Name string `json:name`
	Area uint64	`json:area`
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		// Your resource creation logic goes here. For now it is plain print to console
		fmt.Printf("Got %s city with area of %d sq miles!\n", tempCity.Name, tempCity.Area)
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