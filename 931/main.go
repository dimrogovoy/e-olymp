package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var a, total int
	dob := 1
	for n > 0 {
		a = n % 10
		n = n / 10
		total += a
		dob *= a
	}
	fmt.Printf("%.3f", float64(dob)/float64(total))
}
