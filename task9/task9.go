// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.

package main

import "fmt"

func WriteToChannel(arr ...int) chan int {
	someChannel := make(chan int)
	go func() {
		defer close(someChannel)
		for _, v := range arr {
			someChannel <- v
		}
	}()
	return someChannel
}

func Squaring(someChannel chan int) chan int {
	anotherChannel := make(chan int)
	go func() {
		defer close(anotherChannel)
		for v := range someChannel {
			anotherChannel <- v * v
		}
	}()
	return anotherChannel
}

func main() {
	channelOne := WriteToChannel(1, 2, 3, 4)
	channelTwo := Squaring(channelOne)

	for v := range channelTwo {
		fmt.Println(v)
	}

}
