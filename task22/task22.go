/*
Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b,
значение которых > 2^20.
*/
package main

import (
	"fmt"
	"math/big"
	"math/rand"
)

func addition(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func subtraction(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

func multiplication(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func division(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func main() {
	a := big.NewInt((rand.Int63n(11) + 1) * 1 << (rand.Int63n(31) + 20))
	b := big.NewInt((rand.Int63n(11) + 1) * 1 << (rand.Int63n(31) + 20))
	fmt.Println(a, b)
	fmt.Println(addition(a, b))
	fmt.Println(subtraction(a, b))
	fmt.Println(multiplication(a, b))
	fmt.Println(division(a, b))
}
