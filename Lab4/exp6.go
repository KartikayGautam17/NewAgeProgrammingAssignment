package main

import (
	"fmt"
	"sync"
	"time"
)

type Result struct {
	Task   string
	Values []int
}

func squares(n int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("squares goroutine started")

	var vals []int
	for i := 1; i <= n; i++ {
		vals = append(vals, i*i)
		time.Sleep(30 * time.Millisecond)
	}

	fmt.Println("squares goroutine finished")
	ch <- Result{"Squares", vals}
}

func cubes(n int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("cubes goroutine started")

	var vals []int
	for i := 1; i <= n; i++ {
		vals = append(vals, i*i*i)
		time.Sleep(50 * time.Millisecond)
	}

	fmt.Println("cubes goroutine finished")
	ch <- Result{"Cubes", vals}
}

func fibonacci(n int, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("fibonacci goroutine started")

	var vals []int
	a, b := 0, 1
	for i := 0; i < n; i++ {
		vals = append(vals, a)
		a, b = b, a+b
		time.Sleep(20 * time.Millisecond)
	}

	fmt.Println("fibonacci goroutine finished")
	ch <- Result{"Fibonacci", vals}
}

func main() {
	ch := make(chan Result, 3)
	var wg sync.WaitGroup

	fmt.Println("main: starting goroutines")
	wg.Add(3)
	go squares(10, ch, &wg)
	go cubes(10, ch, &wg)
	go fibonacci(10, ch, &wg)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for r := range ch {
		fmt.Printf("main got %s: %v\n", r.Task, r.Values)
	}

	fmt.Println("main: done")
}
