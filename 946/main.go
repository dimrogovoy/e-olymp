package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func main() {
	slice := make([]Point, 4)
	for i := 0; i < 4; i++ {
		fmt.Scanf("%f %f", &slice[i].x, &slice[i].y)
	}
	var s float64
	s = (slice[0].x*(slice[1].y-slice[3].y) +
		slice[1].x*(slice[2].y-slice[0].y) +
		slice[2].x*(slice[3].y-slice[1].y) +
		slice[3].x*(slice[0].y-slice[2].y)) / 2
	fmt.Println(math.Ceil(math.Abs(s)))
}
