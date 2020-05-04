package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var counter int64

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine(),"\n")

	const gs = 1000
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()
		}()

		fmt.Println("Goroutines:", runtime.NumGoroutine())
	}

	fmt.Println("\nGoroutines:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("count:", counter)
}
