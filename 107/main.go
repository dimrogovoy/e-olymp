package main

import "fmt"

func main() {
	var a, b, c, e, d int
	fmt.Scanf("%d", &d)
	a = (d / 100) * 100
	e = d % 100
	b = e / 20 * 30
	c = d % 20 * 2

	switch {
	case a+b+c > ((d/100)+1)*100:
		fmt.Println(((d / 100) + 1) * 100)
	case a == 0 && a+b+c > ((e/20)+1)*30:
		fmt.Println(((e / 20) + 1) * 30)
	case a != 0 && a+b+c > a+((e/20)+1)*30:
		fmt.Println(a + ((e/20)+1)*30)
	default:
		fmt.Println(a + b + c)
	}

}
