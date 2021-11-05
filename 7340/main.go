package main

import "fmt"

func main() {
	m := make(map[string]bool)
	var s string
	fmt.Scanf("%s", &s)
	for _, v := range s {
		m[string(v)] = true
	}
	fmt.Println(len(m))

}
