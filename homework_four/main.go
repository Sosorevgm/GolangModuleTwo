package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(Workers(0))
}

func Workers(counter int) int {
	var mChannel = make(chan int, 100)
	mChannel <- counter

	for i := 0; i < 1000; i++ {
		go func() {
			numberToIncrement := <-mChannel
			numberToIncrement++
			counter = numberToIncrement
			mChannel <- numberToIncrement
		}()
	}

	time.Sleep(1 * time.Second)
	return counter
}
