package main

import "fmt"

func main() {
	tenge := 50000
	const dollarExchangeRate = 478.5

	result := float64(tenge) / dollarExchangeRate

	fmt.Printf("%d тенге это %.2f долларов\n", tenge, result)
}
