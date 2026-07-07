package main

import "fmt"

func main() {
	var name string = "Edige"

	age := 21

	height := 178.85

	fmt.Printf("Меня зовут %s\nМне %d год\nМой рост: %.2f\n", name, age, height)

	fmt.Printf("Мой рост: %d\n", int(height))
}
