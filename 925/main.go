package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

type Triangle struct {
	a float64
	b float64
	c float64
}

func (t *Triangle) distance(x Point, y Point, z Point) {
	t.a = math.Sqrt((x.x-y.x)*(x.x-y.x) + (x.y-y.y)*(x.y-y.y))
	t.b = math.Sqrt((y.x-z.x)*(y.x-z.x) + (y.y-z.y)*(y.y-z.y))
	t.c = math.Sqrt((z.x-x.x)*(z.x-x.x) + (z.y-x.y)*(z.y-x.y))

}
func (t Triangle) getPerimetr() float64 {
	return t.a + t.b + t.c

}
func (t Triangle) getArea() float64 {
	var p float64
	p = (t.a + t.b + t.c) / 2
	return math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c))
}
func main() {
	var x, y, z Point
	fmt.Scanf("%f %f %f %f %f %f", &x.x, &x.y, &y.x, &y.y, &z.x, &z.y)
	var l Triangle
	l.distance(x, y, z)

	fmt.Printf("%.4f %.4f", l.getPerimetr(), l.getArea())

}
