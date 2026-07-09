package main

import "fmt"

func main() {
	n := 12345

	result := 0

	for n > 0 {
		d := n % 10
		result += d
		n /= 10
	}

	fmt.Println(result)
}
