package main

import (
	"fmt"
	"math"
)

func main() {
	var s, r1 float64
	fmt.Scanf("%f %f", &s, &r1)

	r2 := math.Sqrt(((math.Pi * (r1 * r1)) - s) / math.Pi)
	fmt.Printf("%.2f", r2)
}
