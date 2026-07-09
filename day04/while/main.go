package main

import "fmt"

func main() {
	n := 100

	for n > 1 {
		n /= 2
		fmt.Println(n)
	}
}
