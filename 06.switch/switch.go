package main

import "fmt"

func main() {

	i := 2

	switch i {
		case 1:
			fmt.Println("I'm the one")
		case 2:
			fmt.Println("I'm the beta")
		default:
			fmt.Println("whatever")
	}

}

