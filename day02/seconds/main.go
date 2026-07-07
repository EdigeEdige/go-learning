package main

import "fmt"

func main() {
	inputSeconds := 7385

	hours := inputSeconds / 3600

	minutes := inputSeconds % 3600 / 60

	seconds := inputSeconds % 60

	fmt.Printf("%d seconds = %d hours %d minutes %d seconds\n", inputSeconds, hours, minutes, seconds)
}
