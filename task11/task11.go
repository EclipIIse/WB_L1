// Реализовать пересечение двух неупорядоченных множеств.
package main

import "fmt"

func Intersection(a *[]int, b *[]int) {
	c := make(map[int]struct{})
	var arr []int

	for _, v := range *a {
		c[v] = struct{}{}
	}
	for _, v := range *b {
		if _, ok := c[v]; ok {
			arr = append(arr, v)
		}
	}
	fmt.Println(arr)
}

func main() {
	a := []int{11, 3, 44, 9, 4, 10}
	b := []int{4, 5, 10, 32, 78, 99, 9}

	Intersection(&a, &b)
}
