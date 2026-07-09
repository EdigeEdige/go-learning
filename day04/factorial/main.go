package main

import "fmt"

func main() {
	N := 5

	result := 1

	for i := 1; i <= N; i++ {
		result *= i
	}

	fmt.Println(result)
}
