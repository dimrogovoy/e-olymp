package main

import (
	"fmt"
	"math"
)

type Triangle struct {
	a, b, c float64
}

func (t Triangle) height() (float64, float64, float64) {
	var p, ha, hb, hc float64
	p = (t.a + t.b + t.c) / 2
	ha = (2 / t.a) * math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
	hb = (2 / t.b) * math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
	hc = (2 / t.c) * math.Sqrt(p*(p-t.a)*(p-t.b)*(p-t.c))
	return ha, hb, hc
}

func main() {
	var v Triangle
	fmt.Scanf("%f %f %f", &v.a, &v.b, &v.c)
	a, b, c := v.height()
	fmt.Printf("%.2f %.2f %.2f", a, b, c)
}
