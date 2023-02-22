package main

import (
	"fmt"
	"sync"
)

func AnotherSquaring(num int, wg *sync.WaitGroup) {
	fmt.Println(num * num)
	wg.Done()
}
func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))
	for _, v := range arr {
		go AnotherSquaring(v, wg)
	}
	wg.Wait()
}
