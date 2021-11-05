package main

import "fmt"

type Time struct {
	h, m, s int
}

func main() {
	arr := make([]Time, 2)
	for i := 0; i < 2; i++ {
		fmt.Scanf("%d %d %d", &arr[i].h, &arr[i].m, &arr[i].s)
	}

	var a, b, c int

	if arr[1].s < arr[0].s {
		arr[1].m = arr[1].m - 1
		arr[1].s = arr[1].s + 60
	}
	c = arr[1].s - arr[0].s

	if arr[1].m < arr[0].m {
		arr[1].h = arr[1].h - 1
		arr[1].m = arr[1].m + 60
	}
	b = arr[1].m - arr[0].m
	a = arr[1].h - arr[0].h

	fmt.Println(a, b, c)
}
