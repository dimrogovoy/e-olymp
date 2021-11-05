package main

import "fmt"

func main() {
	var n, a int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		if a >= 0 {
			a += 2
		}
		fmt.Printf("%d ", a)
	}

}
