package main

import "fmt"

func main() {
	var n, b uint
	fmt.Scanf("%d", &n)
	switch {
	case n == 1:
		b = 12
	case n > 1 && n <= 4:
		b = 12 + 8*(n-1)
	case n > 4:
		b = 12 + (8 * 3) + (5 * (n - 4))
	default:
		b = 0
	}

	fmt.Println(b)
}
