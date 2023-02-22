// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
package main

import "fmt"

type Human struct {
	Age  int
	Name string
}

type Action struct {
	Human
}

func (h *Human) getHumanName() {
	fmt.Println(h.Name)
}

func (h *Human) getHumanAge() {
	fmt.Println(h.Age)
}

func main() {
	someHuman := Action{Human{Age: 25, Name: "Koryun"}}
	fmt.Println(someHuman)
	fmt.Println(someHuman.Age)
	someHuman.getHumanName()
	someHuman.getHumanAge()
}
