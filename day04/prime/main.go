package main

import "fmt"

func main() {
	n := 17

	isPrime := true

	for i := 2; i <= n-1; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}
	fmt.Println(isPrime)
}
