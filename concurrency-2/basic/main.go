package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go func ()  {  // thats how we define inner function
		defer wg.Done()  // defer keyword says do this code after all other surrounding code is executed. its easier to do it because we cannot pass wg as a function parameter cause its copied 
		//solution is to pass pointer that way no copy will be passed
		fmt.Println("1")
	} ()

	go func ()  {
		defer wg.Done()
		fmt.Println("2")
	} ()

	go func ()  {
		defer wg.Done() // it is important to add defer no matter what cause order of execution inside func can be changed if parts of inside func is independent of each other
		fmt.Println("3")
	} ()

	wg.Wait()
}