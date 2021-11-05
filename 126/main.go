package main

import "fmt"

func main() {
	var n, p, q, k, skp, skq, nq, np int
	fmt.Scanf("%d %d %d %d", &n, &p, &q, &k)
	skp = n / p
	skq = skp / q
	for j := 1; j <= p && np == 0; j++ {
		switch {
		case k > skp*j && k < skp*(j+1):
			np = j + 1
		case k == skp*j:
			np = j
		}
	}
	if np == 0 {
		np = 1
	}
	var max, min int
	max = skp * np
	min = (max - skp) + 1
	for i := 1; i <= q; i++ {
		if k >= min && k < min+skq {
			nq = i
		}
		min = min + skq

	}
	fmt.Println(np, nq)
}
