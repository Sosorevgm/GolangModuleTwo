package main

import (
	"math/rand"
	"sync"
	"time"
)

var (
	mutex sync.Mutex
)

func main() {
	RunNGoroutines(1000)
	LockMutex()
}

func RunNGoroutines(n int) {
	wg := sync.WaitGroup{}
	wg.Add(n)

	for i := 0; i < n; i++ {
		go func() {
			rand.Seed(time.Now().UnixNano())
			n := rand.Intn(5)
			time.Sleep(time.Duration(n) * time.Second)
			wg.Done()
		}()
	}

	wg.Wait()
}

func LockMutex() {
	mutex.Lock()
	defer mutex.Unlock()
}
