package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("Begin CPU:\t\t\t", runtime.NumCPU())
	fmt.Println("Begin Goroutine:\t\t", runtime.NumGoroutine())
	fmt.Println()

	wg.Add(2)
	go func() {
		fmt.Println("hello from thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from thing two")
		wg.Done()
	}()

	fmt.Println("Mid CPU:\t\t\t", runtime.NumCPU())
	fmt.Println("Mid Goroutine:\t\t\t", runtime.NumGoroutine())
	fmt.Println()

	wg.Wait()
	fmt.Println("about to exit")
	fmt.Println()

	fmt.Println("End CPU:\t\t\t", runtime.NumCPU())
	fmt.Println("End Goroutine:\t\t\t", runtime.NumGoroutine())
}
