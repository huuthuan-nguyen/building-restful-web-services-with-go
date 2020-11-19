package main

import (
	"flag"
	"log"
)

var name string
func init() {
	flag.StringVar(&name, "name", "Kean", "your wonderful name")
}
var age = flag.Int("age", 0, "your graceful age")

func main() {
	flag.Parse()
	log.Printf("Hello %s (%d years), Welcome to the command line world", name, *age)
}