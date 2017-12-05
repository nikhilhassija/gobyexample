package main

import "fmt"
import "math"

func main() {
	var a string = "hello"
	var b string = "world"

	fmt.Println(a + " " + b)
	fmt.Println(a, b)

	var c int = 2
	var d = 3

	fmt.Println(c + d)

	var e int = 2
	var f float64 = 4.0

	fmt.Println(f)
	fmt.Println(float64(e) + f)

	g := 4.0
	h := 2

	h = 2 + h

	fmt.Println(math.Pow(g, 1 / float64(h)))
}