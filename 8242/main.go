package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	switch {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}

}
