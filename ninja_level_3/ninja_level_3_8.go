package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Should not print")
	case true:
		fmt.Println("Should print")
	}
}
