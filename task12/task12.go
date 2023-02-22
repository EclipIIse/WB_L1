// Package task12 Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func main() {

	list := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	for _, v := range list {
		set[v] = struct{}{}
	}

	for v := range set {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
