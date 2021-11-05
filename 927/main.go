package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var a int
	var b float64
	var count int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d %f", &a, &b)
		if b < 50.00 {
			count += a
		}

	}
	fmt.Println(count)
}
