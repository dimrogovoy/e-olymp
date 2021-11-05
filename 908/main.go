package main

import "fmt"

func main() {
	var n, a int
	fmt.Scanf("%d", &n)
	var total, count int
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		if a > 0 && a%6 == 0 {
			count++
			total += a
		}

	}
	fmt.Println(count, total)
}
