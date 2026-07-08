package main

import "fmt"

func main() {
	var grade int

	fmt.Scan(&grade)

	switch {
	case grade < 0 || grade > 100:
		fmt.Println("Ошибка, оценки должны быть в промежутке 0 - 100")
	case grade >= 90:
		fmt.Println("A")
	case grade >= 75:
		fmt.Println("B")
	case grade >= 60:
		fmt.Println("C")
	case grade >= 40:
		fmt.Println("D")
	default:
		fmt.Println("F")
	}
}
