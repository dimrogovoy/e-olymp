package main

import "fmt"

func main() {
	var a, b, c, d byte
	fmt.Scanf("%d %d %d %d", &a, &b, &c, &d)
	m := make(map[byte]struct{})
	if a > b {
		a, b = b, a
	}
	if c > d {
		c, d = d, c
	}
	var i, j byte
	for i = a; i <= b; i++ {
		for j = c; j <= d; j++ {
			m[i*j] = struct{}{}
		}

	}
	fmt.Println(len(m))
}
