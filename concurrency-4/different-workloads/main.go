package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go task1(&wg)
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
	go task5(&wg)

	wg.Wait()

}

func task1 (wg *sync.WaitGroup){
	//http request
	defer wg.Done()
	_, err := http.Get("http://localhost:8080")

	if err != nil {
		log.Fatalf("could not make http request: %v", err)
	}
	fmt.Println("task1 is done")

}

func task2 (wg *sync.WaitGroup){
	//crazy calculation
	defer wg.Done()

	var  count int
	for i:=0; i<100000000; i++{
		count+=i
	}
	fmt.Println("task2 is done")

}

func task3 (wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Println("task3 is done")

}

func task4 (wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(500*time.Millisecond)
	fmt.Println("task4 is done")

}

func task5 (wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(600*time.Millisecond)
	fmt.Println("task5 is done")

}