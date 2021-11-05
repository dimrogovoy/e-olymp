package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scanf("%d %d %d", &a, &b, &c)
	switch {
	case a*a == b*b+c*c:
		fmt.Println("YES")
	case b*b == a*a+c*c:
		fmt.Println("YES")
	case c*c == a*a+b*b:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}

}
