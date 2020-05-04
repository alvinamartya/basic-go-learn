package main

import "fmt"

func main()  {
	s := struct {
		first string
		friends map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q": 777,
			"M": 888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(s)
}
