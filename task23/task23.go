// Удалить i-ый элемент из слайса.
package main

import "fmt"

func removeFromSlice(slice []int, s int) []int {
	result := append(slice[:s], slice[s+1:]...)
	fmt.Println(result)
	return result
}

func main() {
	someSlice := []int{3, 9, 34, 44, 22, 90}
	removeFromSlice(someSlice, 3)
}
