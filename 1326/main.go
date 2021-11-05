package main

import "fmt"

func main() {
	var n, g int
	m := 3
	fmt.Scanf("%d", &n)
	g = n * (n - 1) * (n - m + 1)
	if n == 1 {
		g = 1
	}
	if n == 2 {
		g = 2
	}
	fmt.Println(g)
}
