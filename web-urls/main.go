package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"log"
	"weburls/models"
	base62 "weburls/utils"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// DB stores the database session information. Needs to be initialized once
type DBClient struct {
	db *sql.DB
}

// Model the record struct
type Record struct {
	Id int `json:"id"`
	Url string `json:"url"`
}

// GetOriginalURL fetches the original URL for the given encoded (short) string
func (driver *DBClient) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	var url string
	vars := mux.Vars(r)
	// Get ID from base62 string
	id := base62.ToBase10(vars["encoded_string"])
	err := driver.db.QueryRow("SELECT url FROM web_url WHERE id = $1", id).Scan(&url)
	// Handle response details
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		responseMap := map[string]interface{}{"url": url}
		response, _ := json.Marshal(responseMap)
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

// GenerateShortURL adds URL to DB and gives back shortened string
func (driver *DBClient) GenerateShortURL(w http.ResponseWriter, r *http.Request) {
	var id int
	var record Record
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &record)
	err := driver.db.QueryRow("INSERT INTO web_url(url) VALUES($1) RETURNING id", record.Url).Scan(&id)
	responseMap := map[string]interface{}{"encoded_string": base62.ToBase62(id)}

	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(responseMap)
		w.Write(response)
	}
}

func main() {
	db, err := models.InitDB()
	if err != nil {
		panic(err)
	}
	dbclient := &DBClient{db: db}
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// create a new router
	r := mux.NewRouter()
	// attach an elegant path with handler
	r.HandleFunc("/v1/short/{encoded_string:[a-zA-Z0-9]*}", dbclient.GetOriginalURL).Methods("GET")
	r.HandleFunc("/v1/short", dbclient.GenerateShortURL).Methods("POST")

	server := &http.Server{
		Handler: r,
		Addr: ":8080",
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}

	log.Fatal(server.ListenAndServe())
}