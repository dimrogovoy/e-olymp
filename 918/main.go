package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scanf("%f %f", &x, &y)
	switch {
	case x > 0 && y > 0:
		fmt.Println(1)
	case x < 0 && y > 0:
		fmt.Println(2)
	case x < 0 && y < 0:
		fmt.Println(3)
	case x > 0 && y < 0:
		fmt.Println(4)
	default:
		fmt.Println(0)

	}

}
