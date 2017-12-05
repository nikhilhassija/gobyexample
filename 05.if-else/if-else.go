package main

import "fmt"

func main() {
	for n := 0; n < 10; n++ {
		if n % 2 == 1 {
			fmt.Println(n, "is odd")
		} else {
			fmt.Println(n, "is even")
		}
	}
}
