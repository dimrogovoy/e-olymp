package main

import (
	"fmt"
	"math"
)

//структура яка зберігає точки
type Point struct {
	x, y float64
}

//метод для пошуку координат точки перетину
func (p Point) coordinates(a Point, c Point) (float64, float64) {
	p.x = (a.x + c.x) / 2
	p.y = (a.y + c.y) / 2
	return p.x, p.y
}

//структура яка берегтиме наші діагоналі
type Diagonal struct {
	d1, d2 float64
}

//метод, що вираховує діагоналі
func (d *Diagonal) distance(q Point, w Point, e Point, r Point) (float64, float64) {
	d.d1 = math.Sqrt((e.x-q.x)*(e.x-q.x) + (e.y-q.y)*(e.y-q.y))
	d.d2 = math.Sqrt((r.x-w.x)*(r.x-w.x) + (r.y-w.y)*(r.y-w.y))
	return d.d1, d.d2
}

func main() {
	var a, b, c, d Point
	// вводимо наші координати
	fmt.Scanf("%f %f\n%f %f\n%f %f\n%f %f", &a.x, &a.y, &b.x, &b.y, &c.x, &c.y, &d.x, &d.y)

	var g Diagonal
	d1, d2 := g.distance(a, b, c, d)

	var o Point
	x, y := o.coordinates(a, c)

	fmt.Printf("%.3f %.3f\n%.3f %.3f", x, y, d1, d2)
}
