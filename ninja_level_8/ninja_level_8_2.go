package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string `json:"First"`
	Last    string `json:"Last"`
	Age     int    `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred"]}]`

	var persons []person
	err := json.Unmarshal([]byte(s), &persons)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(persons)

	for i, person := range persons {
		fmt.Println("Person #", i)
		fmt.Println(person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}
