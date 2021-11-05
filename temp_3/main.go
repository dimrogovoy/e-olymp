package main

import (
	"fmt"
	"math/rand"
	"time"
)

// boring is a function that returns a channel to communicate with it.
// <-chan string means receives-only channel of string.
func boring(msg string) <-chan string {
	c := make(chan string)
	// we launch goroutine inside a function
	// that sends the data to channel
	go func() {
		// The for loop simulate the infinite sender.
		for i := 0; i < 10; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}

		// The sender should close the channel
		//close(c)

		fmt.Println(msg, "leaving")

	}()
	return c // return a channel to caller.
}

func Sleep(duration time.Duration) {
	<-time.After(duration)
}

func main() {

	joe := boring("Joe")
	ahn := boring("Ahn")

	// This loop yields 2 channels in sequence
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-joe)
	//	fmt.Println(<-ahn)
	//}

	// or we can use the for loop
	for {
		select {
		case msg := <-joe:
			fmt.Println(msg)
		case msg := <-ahn:
			fmt.Println(msg)
		case <-time.After(2 * time.Second):
			return
		}
	}

	fmt.Println("You're both boring. I'm leaving")
}
