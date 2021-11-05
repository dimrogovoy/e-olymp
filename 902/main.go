package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	switch n {
	case 1, 2, 3:
		fmt.Println("Initial")
	case 4, 5, 6:
		fmt.Println("Average")
	case 7, 8, 9:
		fmt.Println("Sufficient")
	case 10, 11, 12:
		fmt.Println("High")
	}

}
