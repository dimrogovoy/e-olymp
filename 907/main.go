package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d", &n)
	slice := make([]float64, 0, n)
	for i := 0; i < n; i++ {
		var a float64
		fmt.Scanf("%f", &a)
		slice = append(slice, a)
	}
	for i, b := range slice {
		if b <= 2.5 {
			fmt.Printf("%d %.2f", i+1, b)
			return
		}

	}
	fmt.Println("Not Found")
}
