package main

import "fmt"

func main() {
	var age int
	fmt.Scan(&age)

	if age >= 60 {
		fmt.Println("пенсионер")
	} else if age >= 18 {
		fmt.Println("взрослый")
	} else if age >= 13 {
		fmt.Println("подросток")
	} else if age >= 0 {
		fmt.Println("ребёнок")
	} else {
		fmt.Println("Возраст не может быть отрицательным числом")
	}

}
