package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	var a, b int
	a = n / 100
	b = n % 10
	switch {
	case a > b:
		fmt.Println(a)
	case a < b:
		fmt.Println(b)
	case a == b:
		fmt.Println("=")
	}
}
