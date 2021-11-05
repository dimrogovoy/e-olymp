package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Scanf("%d", &a)
	fmt.Println(math.Abs(float64((a / 10) + (a % 10))))

}
