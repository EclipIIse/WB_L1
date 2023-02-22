/* Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/

package main

import "fmt"

func TypeOfVariable(v interface{}) {
	switch v.(type) {
	case string:
		fmt.Println("String")
	case int:
		fmt.Println("Int")
	case bool:
		fmt.Println("Bool")
	case struct{}:
		fmt.Println("Struct")
	case chan int:
		fmt.Println("Channel")
	default:
		fmt.Println("This type is undefined")
	}
}

func main() {
	TypeOfVariable(999)
	TypeOfVariable("Koryun")
	TypeOfVariable(true)
	TypeOfVariable(make(chan int))
	TypeOfVariable(struct{}{})
	TypeOfVariable(13.3)

}
