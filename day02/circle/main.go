package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 5.0

	circleArea := math.Pi * radius * radius

	fmt.Printf("Радиус: %.0f, площадь: %.2f\n", radius, circleArea)
}
