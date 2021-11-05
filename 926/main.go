package main

import (
	"fmt"
	"math"
)

func sTriangle(a, b, c float64) float64 {
	p := (a + b + c) / 2
	s := math.Sqrt(p * (p - a) * (p - b) * (p - c))
	return s
}
func main() {
	var a, b, c, d, f float64
	fmt.Scanf("%f %f %f %f %f", &a, &b, &c, &d, &f)
	var s1, s2 float64
	s1 = sTriangle(c, d, f)
	s2 = sTriangle(a, b, f)
	fmt.Printf("%.4f", s1+s2)

}
