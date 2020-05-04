package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 7, 9, 1, 4, 2}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
