package main

import "fmt"

func main() {
	num := 7

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", num, i, i*num)
	}
}
