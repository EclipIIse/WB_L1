// Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"fmt"
	"sort"
)

func binSearch(arr []int, key int) (int, bool) {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (left + right) / 2
		if arr[middle] > key {
			right = middle - 1
		} else if arr[middle] < key {
			left = middle + 1
		} else {
			return middle, true
		}
	}
	return 0, false
}

func main() {
	arr := []int{34, -200, 0, 11, 12, 1, -33, 99, 4, 3, -4, 13}
	sort.Ints(arr)
	key := 5
	fmt.Printf("Sorted array: %v\n", arr)
	fmt.Println(binSearch(arr, key))
}
