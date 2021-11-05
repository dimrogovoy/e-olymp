package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	switch {
	case a > b && a < c || a < b && a > c:
		fmt.Println(a)
	case b > a && b < c || b < a && b > c:
		fmt.Println(b)
	case c > b && c < a || c < b && c > a:
		fmt.Println(c)

	}

}
