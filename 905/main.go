package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)

	switch {
	case a == b && b == c:
		fmt.Println(1)
	case a == b || a == c || b == c:
		fmt.Println(2)
	case a != b, a != c, b != c:
		fmt.Println(3)

	}

}
