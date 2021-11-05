package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

type Triangle struct {
	a, b, c float64
}

func (t *Triangle) dictance(q Point, w Point, e Point) {
	t.a = math.Sqrt((w.x-q.x)*(w.x-q.x) + (w.y-q.y)*(w.y-q.y))
	t.b = math.Sqrt((e.x-w.x)*(e.x-w.x) + (e.y-w.y)*(e.y-w.y))
	t.c = math.Sqrt((q.x-e.x)*(q.x-e.x) + (q.y-e.y)*(q.y-e.y))

}

func (t Triangle) maxAngle() float64 {
	arr := make([]float64, 3)
	max := .0
	arr[0] = math.Acos((t.a*t.a + t.b*t.b - t.c*t.c) / (2 * t.a * t.b))
	arr[1] = math.Acos((t.b*t.b + t.c*t.c - t.a*t.a) / (2 * t.b * t.c))
	arr[2] = math.Acos((t.c*t.c + t.a*t.a - t.b*t.b) / (2 * t.c * t.a))
	for i := 0; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func main() {
	arr := make([]Point, 3)
	for i := 0; i < len(arr); i++ {
		fmt.Scanf("%f %f", &arr[i].x, &arr[i].y)
	}
	var l Triangle
	l.dictance(arr[0], arr[1], arr[2])
	fmt.Printf("%.6f", l.maxAngle()*(180/math.Pi))
}
