package main

import "fmt"

func main() {
	var n int
	fmt.Scanf("%d", &n)
	//var a,b,c int
	//a=n/100
	//b=n/10%10
	//c=n%10
	//fmt.Println(a*b*c)
	fmt.Println((n / 100) * (n / 10 % 10) * (n % 10))

}
