package main

import (
	"fmt"
)

func main() {
	var n, a int
	fmt.Scanf("%d", &n)
	slice := make([]int, 0, n)
	min := 100
	max := -100
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		slice = append(slice, a)

		if min > slice[i] {
			min = slice[i]
		}
		if max < slice[i] {
			max = slice[i]
		}
	}
	fmt.Println(min + max)
}
