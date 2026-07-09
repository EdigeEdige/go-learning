package main

import "fmt"

func main() {
	i := 1

	counter := 0

	for {
		if counter == 5 {
			break
		}

		if i%2 == 0 {
			i++
			continue
		}

		fmt.Println(i)
		i++
		counter++
	}
}
