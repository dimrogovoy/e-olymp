package main

import "fmt"

func main() {
	var n, s, p int
	for {
		if _, err := fmt.Scanf("%d", &n); err != nil {
			return
		}
		s = n * n
		p = (n + n) * 2
		fmt.Println(p, s)
	}

}
