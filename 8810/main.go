package main

import "fmt"

func main() {
	var a, b, c, total int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	total = a + (b - c)
	fmt.Println(total)

}
