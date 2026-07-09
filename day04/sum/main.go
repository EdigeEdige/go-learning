package main

import "fmt"

func main() {
	N := 10

	sum := 0

	for i := 1; i <= N; i++ {
		sum += i
	}

	fmt.Println(sum)
}
