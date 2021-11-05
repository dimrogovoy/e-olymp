package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var total, m float64
	var count int
	for i := 0; i < n; i++ {
		fmt.Scanf("%f", &m)
		if m > 0 {
			total += m
			count++
		}

	}
	if total == 0 {
		fmt.Println("Not Found")

	} else {
		fmt.Printf("%.2f", total/float64(count))
	}

}
