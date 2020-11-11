package main

import (
	"encoding/json"
	"github.com/justinas/alice"
	"log"
	"net/http"
	"strconv"
	"time"
	"fmt"
)

type city struct {
	Name string `json:name`
	Area uint64	`json:area`
}

// middleware to check content type as JSON
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in the check content type middleware")
		// filtering request by MIME type
		if r.Header.Get("Content-Type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type"))
			return
		}
		// filtering run before the main handler
		handler.ServeHTTP(w, r)
	})
}

// middleware to add server timestamp for response cookie
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// add server timestamp run after main handler
		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

// main handler
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
	mainLogicHandler := http.HandlerFunc(mainHandler)
	chain := alice.New(filterContentType, setServerTimeCookie).Then(mainLogicHandler)
	http.Handle("/city", chain)
	http.ListenAndServe(":8080", nil)
}