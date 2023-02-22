/*
Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
По завершению программа должна выводить итоговое значение счетчика.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	mu    sync.Mutex
	myInt int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.myInt++
	c.mu.Unlock()
}

func (c *Counter) Val() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.myInt
}

func main() {
	c := Counter{myInt: 0}
	for i := 0; i < 100; i++ {
		go c.Inc()
	}
	time.Sleep(time.Second)
	fmt.Println(c.Val())
}
