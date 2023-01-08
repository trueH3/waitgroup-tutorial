package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{}) // channels need to be declared before use, way to do it is to use make keyword
	now := time.Now()
	 go task1(done)
	 go task2(done)

	<-done // this syntax reads value passed in channel, it waits until value will be accessible
	<-done // this syntax reads value passed in channel, it waits until value will be accessible
	fmt.Printf("ovveral time %v\n", time.Since(now))
}


func task1(c chan struct{}) {
	time.Sleep(50*time.Microsecond)
	fmt.Println("task1")
	c <- struct{}{} // this syntax assign to c channel value struct{}{}
}

func task2(c chan struct{}) {
	time.Sleep(100*time.Microsecond)
	fmt.Println("task2")
	c <- struct{}{} // this syntax assign to c channel value struct{}{}
}