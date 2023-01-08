package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	go task1()
	go task2()
	go task3()
	go task4()

	fmt.Printf("ovveral time %v\n", time.Since(now))
}


func task1() {
	time.Sleep(50*time.Microsecond)
	fmt.Println("task1")
}

func task2() {
	time.Sleep(100*time.Microsecond)
	fmt.Println("task2")
}

func task3() {
	time.Sleep(120*time.Microsecond)
	fmt.Println("task3")
}

func task4() {
	time.Sleep(200*time.Microsecond)
	fmt.Println("task4")
}