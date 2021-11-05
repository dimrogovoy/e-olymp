package main

import "fmt"

func SumDub(a float64, b float64) (float64, float64) {
	s := a + b
	d := a * b
	return s, d
}
func main() {
	var n int
	fmt.Scanf("%d", &n)
	var a, b float64

	for i := 0; i < n; i++ {
		fmt.Scanf("%f %f", &a, &b)
		var s, d float64
		s, d = SumDub(a, b)
		fmt.Printf("%.4f %.4f\n", s, d)
	}

}
