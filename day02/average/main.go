package main

import "fmt"

func main() {
	n1 := 7
	n2 := 4
	n3 := 6

	average := float64(n1+n2+n3) / 3

	fmt.Printf("Average of %d, %d, %d is %.2f\n", n1, n2, n3, average)
}
