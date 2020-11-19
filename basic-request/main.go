package main

import (
	"github.com/levigross/grequests"
	"log"
)

func main() {
	resp, err := grequests.Get("http://httpbin.org/get", nil)
	// you can modify the request by passing an optinal RequestOptions struct
	if err != nil {
		log.Fatalln("Unable to make request:", err)
	}
	log.Println(resp.String())
}