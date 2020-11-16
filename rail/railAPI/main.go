package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"
	restful "github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
	"rail/dbutils"
)

// DB driver visible to whole program
var DB *sql.DB

// TrainResource is the model for holding rail information
type TrainResource struct {
	Id int `json:"id,omitempty"`
	DriverName string `json:"driver_name"`
	OperatingStatus bool `json:"operating_status"`
}

// StationResource hold information about locations
type StationResource struct {
	Id int
	Name string
	OpeningTime time.Time
	ClosingTime time.Time
}

// ScheduleResource links both trains and stations
type ScheduleResource struct {
	Id int
	TrainId int
	StationId int
	ArrivalTime time.Time
}

func (t *TrainResource) Register (container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/v1/trains").Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	ws.Route(ws.GET("/{trainId}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{trainId}").To(t.removeTrain))
	container.Add(ws)
}

// GET a specific train
func (t *TrainResource) getTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("trainId")
	err := DB.QueryRow("SELECT id, driver_name, operating_status FROM train WHERE id=?", id).Scan(&t.Id, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println(err)
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusNotFound, "Train could not be found.")
	} else {
		response.WriteEntity(t)
	}
}

// POST to create new train
func (t *TrainResource) createTrain(request *restful.Request, response *restful.Response) {
	log.Println(request.Request.Body)
	decoder := json.NewDecoder(request.Request.Body)
	var b TrainResource
	err := decoder.Decode(&b)
	log.Println(b.DriverName, b.OperatingStatus)
	// error handling is obvious here. So omitting...
	statement, _ := DB.Prepare("INSERT INTO train(driver_name, operating_status) VALUES(?, ?)")
	result, err := statement.Exec(b.DriverName, b.OperatingStatus)
	if err == nil {
		newId, _ := result.LastInsertId()
		b.Id = int(newId)
		response.WriteHeaderAndEntity(http.StatusCreated, b)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

// DELETE a specific train
func (t *TrainResource) removeTrain(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("trainId")
	statement, _ := DB.Prepare("DELETE FROM train WHERE id=?")
	_, err := statement.Exec(id)
	if err == nil {
		response.WriteHeader(http.StatusOK)
	} else {
		response.AddHeader("Content-Type", "text/plain")
		response.WriteErrorString(http.StatusInternalServerError, err.Error())
	}
}

func main() {
	// Connect to Database
	var err error
	DB, err = sql.Open("sqlite3", "./rail.db")

	if err != nil {
		log.Println("Driver creation failed!")
	}
	// create tables
	dbutils.Initialize(DB)
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	t := TrainResource{}
	t.Register(wsContainer)
	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}