package main

import "fmt"

func main() {
	var n uint
	fmt.Scanf("%d", &n)
	var count uint
	if n == 0 {
		count = 1
	}
	for n > 0 {
		n = n / 10
		count++
	}
	fmt.Println(count)
}
