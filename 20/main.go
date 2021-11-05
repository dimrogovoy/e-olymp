package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var count, a int
	for n > 0 {
		tmpn := n
		var total int
		for tmpn > 0 {
			a = tmpn % 10
			tmpn = tmpn / 10
			total += a

		}
		n = n - total
		count++
	}
	fmt.Print(count)
}
