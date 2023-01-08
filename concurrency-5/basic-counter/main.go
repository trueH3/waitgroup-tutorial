package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// to check if there is no race condition just go go run -race ./
func main() {
	var count int64
	var wg sync.WaitGroup
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		time.Sleep(10*time.Millisecond)
		fmt.Println("go routine count =", atomic.LoadInt64(&count))
	}()

	wg.Wait()

	wg.Add(50)
	for i:=0; i<50; i++ {
		go func ()  {
			defer wg.Done()
			time.Sleep(10*time.Millisecond)
			atomic.AddInt64(&count, 1)
		}()
	}

	wg.Wait()
	fmt.Println("count in main:",count)
}