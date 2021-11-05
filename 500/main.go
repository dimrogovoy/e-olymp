package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var l, w, h, b int
	var s float64
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %d %d", &l, &w, &h)
		s = float64((w*h + l*h) * 2)
		b = int(math.Ceil(s / 16))
		fmt.Println(b)
	}

}
