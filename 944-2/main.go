package main

import (
	"fmt"
	"math"
)

func distance(x Point, y Point) float64 {
	return math.Sqrt((y.x-x.x)*(y.x-x.x) + (y.y-x.y)*(y.y-x.y) + (y.z-x.z)*(y.z-x.z))
}

type Point struct {
	x, y, z float64
}

type Triangle struct {
	ab, bc, ca float64
}

func NewTriangle(x Point, y Point, z Point) (*Triangle, error) {
	var t = &Triangle{
		ab: distance(x, y),
		bc: distance(x, z),
		ca: distance(y, z),
	}

	return t, nil
}

func (t Triangle) area() float64 {
	var pp float64
	pp = (t.bc + t.ab + t.ca) / 2
	return math.Sqrt(pp * (pp - t.ab) * (pp - t.bc) * (pp - t.ca))
}

func (t Triangle) perimeter() float64 {
	return t.ab + t.bc + t.ca
}

func (t Triangle) height() float64 {
	return math.Sqrt(t.bc*t.bc - t.ab*t.ab/2)
}

type Square struct {
	ab, bc, cd, da float64
}

func NewSquare(a Point, b Point, c Point, d Point) (*Square, error) {
	var s = &Square{
		ab: distance(a, b),
		bc: distance(b, c),
		cd: distance(c, d),
		da: distance(d, a),
	}

	return s, nil
}

func (s Square) area() float64 {
	return s.ab * s.bc
}

func (s Square) perimeter() float64 {
	return s.ab + s.bc + s.cd + s.da
}

func (s Square) diagonal() float64 {
	return math.Sqrt2 * s.ab
}

type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(sAreas ...Shape) float64 {
	var total float64
	for _, s := range sAreas {
		total += s.area()
	}
	return total
}

func totalPerimeter(shapes ...Shape) float64 {
	var total float64
	for _, s := range shapes {
		total += s.perimeter()
	}
	return total

}

func main() {
	arr := make([]Point, 5)
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%f %f %f", &arr[i].x, &arr[i].y, &arr[i].z)
	}
	var t1, t2, t3, t4 *Triangle
	t1, _ = NewTriangle(arr[0], arr[1], arr[4])
	t3, _ = NewTriangle(arr[2], arr[3], arr[4])
	t2, _ = NewTriangle(arr[1], arr[2], arr[4])
	t4, _ = NewTriangle(arr[3], arr[0], arr[4])
	var p *Square
	p, _ = NewSquare(arr[0], arr[1], arr[2], arr[3])

	s := []Shape{t1, t2, t3, t4, p}
	fmt.Printf("%.3f\n", totalArea(s...))
	fmt.Printf("%.3f", totalPerimeter(s...))

	for _, shape := range s {
		switch v := shape.(type) {
		case Square:
			v.diagonal()
		case Triangle:
			v.height()
		}
	}
}
