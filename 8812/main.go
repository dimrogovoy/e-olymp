package main

import "fmt"

type pryamo struct {
	a int
	b int
}

func (p pryamo) perimetr() int {
	return 2 * (p.a + p.b)
}
func (p pryamo) plosh() int {
	return p.a * p.b
}

func main() {
	var c pryamo
	fmt.Scanf("%d %d", &c.a, &c.b)
	fmt.Println(c.perimetr(), c.plosh())
}
