package main

import (
	"fmt"
	"math"
)

func main() {
	radius := 5.0

	circleRadius := math.Pi * radius * radius

	fmt.Printf("Радиус: %.f, площадь: %.2f\n", radius, circleRadius)
}
