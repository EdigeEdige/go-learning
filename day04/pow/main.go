package main

import "fmt"

func main() {
	base := 2
	exp := 10

	result := 1

	for i := 1; i <= exp; i++ {
		result = base * result
	}

	fmt.Println(result)
}
