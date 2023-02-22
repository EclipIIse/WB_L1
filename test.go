package main

import (
	"fmt"
	"strings"
)

//func update(p *int) {
//	b := 2
//	p = &b
//}
//
//func main() {
//	var (
//		a = 1
//		p = &a
//	)
//	fmt.Println(*p)
//	update(p)
//	fmt.Println(*p)
//}

//func main() {
//	wg := sync.WaitGroup{}
//	for i := 0; i < 5; i++ {
//		wg.Add(1)
//		go func(wg sync.WaitGroup, i int) {
//			fmt.Println(i)
//			wg.Done()
//		}(wg, i)
//	}
//	wg.Wait()
//	fmt.Println("exit")
//}

//func main() {
//	n := 0
//	if true {
//		n := 1
//		n++
//	}
//	fmt.Println(n)
//}

//func someAction(v []int8, b int8) {
//	//v[0] = 100
//	v = append(v, b)
//}
//
//func main() {
//	var a = []int8{1, 2, 3, 4, 5}
//	someAction(a, 6)
//	fmt.Println(a)
//}

//func main() {
//	slice := []string{"a", "a"}
//
//	func(slice []string) {
//		slice = append(slice, "a")
//		slice[0] = "b"
//		slice[1] = "b"
//		fmt.Print(slice)
//	}(slice)
//	fmt.Print(slice)
//}

func concat2builder(x, y string) string {
	var builder strings.Builder
	builder.Grow(len(x) + len(y)) // Только эта строка выделяет память
	builder.WriteString(x)
	builder.WriteString(y)
	return builder.String()
}
func main() {
	fmt.Println(concat2builder("a", "b"))
}
