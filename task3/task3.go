//Дана последовательность чисел: 2,4,6,8,10.
//Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

package main

import (
	"fmt"
	"sync"
)

func Squaring(arr *[5]int, ch chan int, wg *sync.WaitGroup) {
	for _, v := range arr {
		ch <- v * v
	}
	wg.Done()
}
func Summation(ch chan int, lenArr int, wg *sync.WaitGroup) {
	sum := 0
	for i := 0; i < lenArr; i++ {
		sum += <-ch
	}
	fmt.Println(sum)
	wg.Done()
}

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go Squaring(&arr, ch, wg)
	go Summation(ch, len(arr), wg)
	wg.Wait()
}
