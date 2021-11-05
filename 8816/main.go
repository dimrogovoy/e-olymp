package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scanf("%f\n%f", &a, &b)
	c := math.Pow(a, b)

	fmt.Printf("%.0f", c)
}
