package main

import (
	"fmt"
)

func main() {
	var t, k int
	fmt.Scanf("%d", &t)

	for i := 0; i < t; i++ {
		fmt.Scanf("%d", &k)
		var G, C, V string
		G = "G"
		C = "C"
		V = "V"

		for i := 0; i < k; i++ {
			C, V = V, C
			G, C = C, G
		}
		fmt.Println(G + C + V)
	}

}
