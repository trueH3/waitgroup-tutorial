package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	// to test for race condition, run this example with the -race flag
	// go run -race main.go
	var count int32
	var wg sync.WaitGroup
	wg.Add(5)
	go func ()  {
		defer wg.Done()
		atomic.AddInt32(&count, 10)
	}()

	go func ()  {
		defer wg.Done()
		atomic.AddInt32(&count, -15)
	}()

	go func ()  {
		defer wg.Done()
		atomic.AddInt32(&count, 1)
	}()

	go func ()  {
		defer wg.Done()
		atomic.StoreInt32(&count, 0)
	}()

	go func ()  {
		defer wg.Done()
		atomic.StoreInt32(&count, 100)

	}()
	wg.Wait()
	fmt.Println("count =", count)
}