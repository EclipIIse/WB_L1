//Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	// Остановка горутины при каком то условии
	num := 0
	wg.Add(1)
	go func() {
		for num == 0 {
			fmt.Printf("Num is %d\n", num)
		}
		fmt.Println("Stopped")
		wg.Done()
	}()
	num++
	wg.Wait()

	// Остановка горутины через контекст
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context) {
		num2 := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context stop..")
				wg.Done()
				return
			default:
				num2++
			}
		}
	}(ctx)
	cancel()
	wg.Wait()

	// Остановка горутины, читающая данные из канала, при закрытии которого горутина остановится
	ch := make(chan int)
	wg.Add(1)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}()
	for i := 0; i < 5; i++ {
		ch <- i
	}
	fmt.Println("Chanel closed..")
	close(ch)
	wg.Wait()

	// Остановка горутины при получении соответствующего сообщения из канала
	ch2 := make(chan string)
	wg.Add(1)
	go func() {
		i := 0
		for {
			select {
			case <-ch2:
				fmt.Println("STOP")
				wg.Done()
				return
			default:
				i++
			}
		}
	}()
	ch2 <- "Close channel"
	wg.Wait()
}
