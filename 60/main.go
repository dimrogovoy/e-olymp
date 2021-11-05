package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func main() {
	var n int
	fmt.Scanf("%d", &n)
	arr := make([]Point, n+1)
	for i := 0; i < n; i++ {
		fmt.Scanf("%f %f", &arr[i].x, &arr[i].y)
	}
	arr[n] = arr[0]
	var s, total, result float64
	for i := 0; i < n; i++ {
		s = (arr[i].x * arr[i+1].y) - (arr[i+1].x * arr[i].y)
		total += s

	}
	result = math.Abs(total) / 2
	fmt.Printf("%.3f", result)
}
