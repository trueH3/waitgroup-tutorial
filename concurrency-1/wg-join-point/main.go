package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg  sync.WaitGroup
	wg.Add(2)
	now := time.Now()
	 go task1(&wg)
	 go task2(&wg)

	wg.Wait()
	fmt.Printf("ovveral time %v\n", time.Since(now))
}


func task1(wg *sync.WaitGroup) { // we need to pass pointer cause otherwise wg would be copied and allocated in different place in memory 
	defer wg.Done() // it is important to add defer no matter what cause order of execution inside func can be changed if parts of inside func is independent of each other
	time.Sleep(50*time.Microsecond)
	fmt.Println("task1")
	
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(100*time.Microsecond)
	fmt.Println("task2")
}