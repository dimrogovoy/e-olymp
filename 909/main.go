package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var b int
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		a := scanner.Text()
		for _, j := range a {
			if string(j) == " " {
				b++
			}
		}
	}

	fmt.Println(b + 1)
}
