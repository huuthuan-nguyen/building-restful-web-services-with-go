package main

import (
	"fmt"
)

func main() {
	increment := generator()
	for i := 0; i < 5; i++ {
		fmt.Print(increment(), "\t")
	}
}

// this function returns another function
func generator() func() int {
	var i = 0
	return func() int {
		i++
		return i
	}
}