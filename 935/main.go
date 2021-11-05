package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Scanf("%d", &n)
	var a, b, c int
	a = int(math.Abs(float64(n / 100)))
	b = int(math.Abs(float64(n / 10 % 10)))
	c = int(math.Abs(float64(n % 10)))
	fmt.Printf("%d\n%d\n%d\n", a, b, c)

}
