package main

import (
	"fmt"
	"math"
)

type Piramida struct {
	d float64
	p float64
}

func (p Piramida) sPiramida() float64 {
	s := ((p.d*math.Sqrt(p.p*p.p-(p.d*p.d)/4))/2)*4 + p.d*p.d
	return s
}

func (p Piramida) vPiramida() float64 {
	v := p.d * p.d / 3 * math.Sqrt(p.p*p.p-(p.d*p.d)/2)
	return v
}

func main() {
	var p Piramida
	fmt.Scanf("%f %f", &p.d, &p.p)
	var s, v float64
	s = p.sPiramida()
	v = p.vPiramida()

	fmt.Printf("%.3f %.3f", s, v)

}
