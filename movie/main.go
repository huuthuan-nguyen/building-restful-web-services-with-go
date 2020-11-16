package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
	"log"
	"github.com/gorilla/mux"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type DB struct {
	session *mgo.Session
	collection *mgo.Collection
}

// movie holds a move data
type Movie struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Name string `json:"name" bson:"name"`
	Year string `json:"year" bson:"year"`
	Directors []string `json:"directors" bson:"directors"`
	Writers []string `json:"writers" bson:"writers"`
	BoxOffice `json:"box_office" bson:"boxOffice"`
}

// box office is nested in Movie
type BoxOffice struct {
	Budget uint64 `json:"budget" bson:"budget"`
	Gross uint64 `json:"gross" bson:"gross"`
}


// GET a specific movie
func (db *DB) GetMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var movie Movie
	err := db.collection.Find(bson.M{"_id": bson.ObjectIdHex(vars["id"])}).One(&movie)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}

// POST a new movie
func (db *DB) PostMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	postBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(postBody, &movie)
	// create hash ID to insert
	movie.Id = bson.NewObjectId()
	err := db.collection.Insert(movie)
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(movie)
		w.Write(response)
	}
}

// PUT to update a movie
func (db *DB) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var movie Movie
	putBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(putBody, &movie)
	// create an Hash ID to insert
	err := db.collection.Update(bson.M{"_id": bson.ObjectIdHex(vars["id"])}, bson.M{"$set": &movie})
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Updated succesfully!"))
	}
}

// DELETE to remove a movie
func (db *DB) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// create an hash id to remove
	err := db.collection.Remove(bson.M{"_id": bson.ObjectIdHex(vars["id"])})
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "text")
		w.Write([]byte("Deleted successfully!"))
	}
}

func main() {
	session, err := mgo.Dial("127.0.0.1")
	c := session.DB("appdb").C("movies")
	db := &DB{session: session, collection: c}
	if err != nil {
		panic(err)
	}

	defer session.Close()

	// create a new router
	r := mux.NewRouter()
	// attach an elegant path with handler
	r.HandleFunc("/v1/movies/{id:[0-9a-zA-Z]+}", db.GetMovie).Methods("GET")
	r.HandleFunc("/v1/movies", db.PostMovie).Methods("POST")
	r.HandleFunc("/v1/movies/{id:[0-9a-zA-Z]+}", db.UpdateMovie).Methods("PUT")
	r.HandleFunc("/v1/movies/{id:[0-9a-zA-Z]+}", db.DeleteMovie).Methods("DELETE")

	server := &http.Server{
		Handler: r,
		Addr: ":8080",
		ReadTimeout: 15*time.Second,
		WriteTimeout: 15*time.Second,
	}

	log.Fatal(server.ListenAndServe())
}