package main

import (
	"fmt"
	"math"
)

func main() {
	var h, w, v, l float64
	var k, n int
	fmt.Scanf("%f %f %f %d", &h, &w, &l, &k)
	v = h * w * l
	n = int(math.Ceil(v / float64(k)))

	fmt.Println(n)
}
