package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)

	var count int
	var total float64
	var m float64
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &m)
		if m < 0 {
			count++
			total += m
		}
	}
	fmt.Printf("%d %.2f", count, total)
}
