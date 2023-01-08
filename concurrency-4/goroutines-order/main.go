package main

import (
	"fmt"
	"sync"
)

// we do not have power to execute go routines in order written in our code 
func main() {
	var wg sync.WaitGroup
	wg.Add(10)

	for i:=0; i<10; i++{
		go func (i int)  {
			defer wg.Done()
			fmt.Println("goroutine", i+1)
		}(i)

	}

	wg.Wait()
}

// but what if I want to preserve order?