package main

import "fmt"

func main() {
	var n uint
	fmt.Scanf("%d", &n)
	var a, b, c, d uint
	a = n / 86400
	n = n - a*86400
	b = n / 60 / 60
	n = n - b*60*60
	c = n / 60
	n = n - c*60
	d = n

	fmt.Println(a, b, c, d)
}
