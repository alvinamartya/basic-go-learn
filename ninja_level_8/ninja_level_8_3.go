package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string
	Last string
	Age int
	Sayings []string
}
func main()  {
	u1 := user{
		First: "James",
		Last: "Bond",
		Age: 32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
		},
	}

	u2 := user{
		First: "M",
		Last: "Hmmm",
		Age: 54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
		},
	}

	users := []user{u1,u2}
	fmt.Println(users)

	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println("We did something wrong and here's the error:",err)
	}
}
