package main

import (
	"os/exec"
	"github.com/julienschmidt/httprouter"
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// exec a system command and return output
func getCommandOutput(command string, arguments ...string) string {
	// args... unpacks arguments array into elements
	cmd := exec.Command(command, arguments...)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Start()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(fmt.Sprint(err) + ": " + stderr.String())
	}
	return out.String()
}

// get Go version
func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("go", "version"))
}

// echo ":name"
func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Fprintf(w, getCommandOutput("echo", params.ByName("name")))
}

func main() {
	// use httprouter as multiplexer
	router := httprouter.New()
	// mapping method
	router.GET("/api/v1/go-version", goVersion)
	// path variable
	router.GET("/api/v1/show-file/:name", getFileContent)
	// serve static file
	router.ServeFiles("/static/*filepath", http.Dir("./static"))

	log.Fatal(http.ListenAndServe(":8080", router))
}