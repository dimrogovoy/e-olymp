package main

import "fmt"

type Parallelogram struct {
	a, b, c, d float64
}

func (p Parallelogram) comparison() string {
	switch {
	case p.a == p.c && p.b == p.d:
		return "YES"
	case p.a == p.b && p.c == p.d:
		return "YES"
	case p.a == p.d && p.b == p.c:
		return "YES"
	default:
		return "NO"
	}
}
func main() {
	var k Parallelogram
	fmt.Scanf("%f %f %f %f", &k.a, &k.b, &k.c, &k.d)
	fmt.Println(k.comparison())
}
