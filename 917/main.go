package main

import (
	"fmt"
)

func main() {
	var n int
	var m float64
	fmt.Scanf("%d", &n)
	var min float64
	min = 100
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &m)
		if min > m {
			min = m
		}
	}
	fmt.Printf("%.2f", min*2)
}
