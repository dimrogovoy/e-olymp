package main

import "fmt"

func main() {
	var a, m int
	fmt.Scanf("%d %d", &a, &m)
	var count int
	for m > 0 {
		m = m - a
		a = a + 1
		count++
		if m == a*2 {
			count += 1
			break
		}
	}
	fmt.Println(count)
}
