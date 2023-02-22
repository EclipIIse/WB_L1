package main

import "fmt"

func AnotherIntersection(a *[]int, b *[]int) {
	var c []int
	for _, v := range *a {
		for _, vv := range *b {
			if v == vv {
				c = append(c, vv)
			}
		}
	}
	fmt.Printf("%v\n", c)
}

func main() {
	a := []int{11, 3, 44, 9, 4, 10}
	b := []int{4, 5, 10, 32, 78, 99, 9}

	AnotherIntersection(&a, &b)
}
