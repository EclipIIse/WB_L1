//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

func BitToOne(n int64, i int) int64 {
	return n | 1<<(i-1)
}

func BitToZero(n int64, i int) int64 {
	return n & ^(1 << (i - 1))
}

func main() {

	number := int64(7)

	test := BitToOne(number, 1)
	test2 := BitToZero(number, 3)

	fmt.Printf("The original number is %d (%b)\n", number, number)
	fmt.Printf("%b\n", test)
	fmt.Printf("%b\n", test2)

}
