package main

import "fmt"

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &arr[i][j])
		}
	}
	arr1 := make([][]int, m)
	for i := 0; i < m; i++ {
		arr1[i] = make([]int, n)
		for j := 0; j < n; j++ {
			arr1[i][j] = arr[j][i]
		}

	}
	for _, j := range arr1 {
		for m, i := range j {
			if m == 0 {
				fmt.Print(i)
			} else {
				fmt.Print(" ", i)

			}

		}
		fmt.Println()
	}
}
