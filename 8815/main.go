package main

import "fmt"

type cube struct {
	a int
}

func (c cube) square() int {
	return (c.a * c.a) * 6
}
func (c cube) volume() int {
	return c.a * c.a * c.a
}
func main() {
	var b cube
	fmt.Scanf("%d", &b.a)
	fmt.Println(b.square(), b.volume())
}
