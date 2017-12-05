package main

import (
	"fmt"
)

var a string = "hello"

func main() {
	b := "world"

	fmt.Println(a, b)

	a := "bye"

	fmt.Println(a, b)
}