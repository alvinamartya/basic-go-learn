package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "skiing":
		fmt.Println("Go to the mountain!")
	case "swimming":
		fmt.Println("Go to the pool!")
	case "surfing":
		fmt.Println("Go to hawaii!")
	case "wing-suit flying":
		fmt.Println("What would you like me to say at your funeral?")
	}
}
