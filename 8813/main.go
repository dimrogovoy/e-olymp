package main

import "fmt"

type Parallelepiped struct {
	a, b, c int
}

func (p Parallelepiped) s() int {
	s := 2 * (p.a*p.b + p.b*p.c + p.a*p.c)
	return s
}

func (p Parallelepiped) p() int {
	v := p.a * p.b * p.c
	return v
}

func main() {
	var z Parallelepiped
	fmt.Scanf("%d %d %d", &z.a, &z.b, &z.c)
	fmt.Println(z.s(), z.p())
}
