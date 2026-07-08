package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Scan(&n1, &n2, &n3)

	valueOfMax := max(n1, n2, n3)

	fmt.Printf("%d is the max number\n", valueOfMax)

	var maxValue int

	if n1 >= n2 && n1 >= n3 {
		maxValue = n1
	} else if n2 >= n1 && n2 >= n3 {
		maxValue = n2
	} else {
		maxValue = n3
	}

	fmt.Printf("Max is %d\n", maxValue)
}
