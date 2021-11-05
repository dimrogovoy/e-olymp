package main

import "fmt"

type kvadrat struct {
	a int
}

func (k kvadrat) perimetr() int {
	return 4 * k.a
}

func (k kvadrat) plosh() int {
	return k.a * k.a
}

func main() {
	var b kvadrat
	fmt.Scanf("%d", &b.a)
	fmt.Println(b.perimetr(), b.plosh())
}
