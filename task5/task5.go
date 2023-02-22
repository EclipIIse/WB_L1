//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ReadFromChannel(ch chan interface{}) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan interface{})

	start := time.Now()
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*2))
	defer cancel()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	go ReadFromChannel(ch)

loop:
	for {
		select {
		case <-shutdown:
			close(ch)
			fmt.Println("See you..")
			break loop
		case <-ctx.Done():
			fmt.Printf("Closed after %v seconds\n", time.Since(start))
			break loop
		default:
			ch <- "Hey there"
			ch <- 25
			ch <- true
			time.Sleep(time.Millisecond * 300)
		}
	}
}
