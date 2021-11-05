package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y, z float64
}

func (p Point) distance(a Point, b Point) float64 {
	d := math.Sqrt((b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y) + (b.z-a.z)*(b.z-a.z))
	return d
}

func main() {
	var a, b, d Point
	fmt.Scanf("%f %f\n%f %f %f", &a.x, &a.y, &b.x, &b.y, &b.z)
	var c float64
	c = 1 / d.distance(a, b)
	fmt.Printf("%.3f", c)

}
