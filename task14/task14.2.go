package main

import (
	"fmt"
	"reflect"
)

func TypeOfVariable2(v interface{}) {
	fmt.Println(reflect.TypeOf(v))
}

func main() {
	TypeOfVariable2(999)
	TypeOfVariable2("Koryun")
	TypeOfVariable2(true)
	TypeOfVariable2(make(chan int))
	TypeOfVariable2(struct{}{})
	TypeOfVariable2(13.3)

}
