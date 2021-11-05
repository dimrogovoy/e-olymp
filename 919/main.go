package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var count int
	var total float64
	slice := make([]float64, 0, n)
	var b float64
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &b)
		slice = append(slice, b)
	}
	for i, v := range slice {
		if (i+1)%3 == 0 && v > 0 {
			count += 1
			total += v
		}
	}
	fmt.Printf("%d %.2f", count, total)

}
