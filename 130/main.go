package main

import (
	"fmt"
)

type Point struct {
	x, y float64
}

func ggg(a Point, b Point) float64 {
	return (b.x-a.x)*(b.x-a.x) + (b.y-a.y)*(b.y-a.y)
}

func sss(a Point, b Point) Point {
	var s Point
	s.x = (a.x + b.x) / 2
	s.y = (a.y + b.y) / 2
	return s
}

func main() {
	var a, b, c Point
	fmt.Scanf("%f %f %f %f %f %f", &a.x, &a.y, &b.x, &b.y, &c.x, &c.y)

	var g1, g2, g3 float64
	g1 = ggg(a, b)
	g2 = ggg(a, c)
	g3 = ggg(b, c)

	var p, s, d Point
	switch {
	case g1 == g2+g3:
		p = c
		s = sss(a, b)
	case g2 == g3+g1:
		p = b
		s = sss(a, c)
	case g3 == g2+g1:
		p = a
		s = sss(b, c)
	}
	d.x = (s.x - (p.x / 2)) * 2
	d.y = (s.y - (p.y / 2)) * 2

	fmt.Println(d.x, d.y)

}
