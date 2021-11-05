package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var m, max float64
	max = 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &m)
		if max < math.Abs(m) {
			max = math.Abs(m)

		}

	}
	fmt.Printf("%.2f", max)
}
