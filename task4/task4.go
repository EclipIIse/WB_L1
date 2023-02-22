//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Worker(ch chan interface{}, count int) {
	for i := range ch {
		fmt.Printf("Worker %v -> %v\n", count, i)
	}
}

func main() {
	workersCount := flag.Int("w", 2, "Workers count")
	flag.Parse()

	//workersCount := 5
	ch := make(chan interface{})

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT)

	for i := 0; i < *workersCount; i++ {
		go Worker(ch, i+1)
	}

loop:
	for {
		select {
		case <-shutdown:
			close(ch)
			fmt.Println("See you..")
			break loop
		default:
			ch <- "Hey there"
			ch <- 25
			ch <- true
			time.Sleep(time.Millisecond * 300)
		}
	}
}
