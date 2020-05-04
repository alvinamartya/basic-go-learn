package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100

	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			fmt.Printf("begin $%d:%d\n", i, incrementer)
			v := incrementer
			runtime.Gosched()
			v++
			incrementer = v
			fmt.Printf("end $%d:%d\n", i, incrementer)
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("end value:", incrementer)
}
