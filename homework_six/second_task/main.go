package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for i := 0; ; i++ {
			fmt.Printf("Goroutine one: %v\n", i)
			if i%100 == 0 {
				fmt.Println("Time to take a break")
				runtime.Gosched()
			}
		}
		wg.Done()
	}()

	time.Sleep(1 * time.Second)

	go func() {
		for i := 0; ; i++ {
			fmt.Printf("Goroutine two: %v\n", i)
			if i%100 == 0 {
				fmt.Println("Time to take a break")
				runtime.Gosched()
			}
		}
	}()
	wg.Wait()
}
