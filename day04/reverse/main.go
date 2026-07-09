package main

import (
	"fmt"
)

func main() {
	n := 12345

	reversed := 0

	for n > 0 {
		d := n % 10

		reversed = reversed*10 + d

		n /= 10
	}

	fmt.Println(reversed)
}
