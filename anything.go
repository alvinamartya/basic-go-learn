package main

import "fmt"

func main() {
	fmt.Println("Hello everyone, this is the most awesome class in the entire world for having")
	foo()
	fmt.Println("Something more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	bar()
}

func bar() {
	fmt.Println("and then we exited")
}

func foo() {
	fmt.Println("I'm in foo")
}
