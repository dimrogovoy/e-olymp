package main

import "fmt"

func main() {
	var n, l, count int
	fmt.Scanf("%d", &n)
	for i := 1; i <= n; i++ {
		l = i * 2
		count += l
	}
	fmt.Println(count + 1)
}
