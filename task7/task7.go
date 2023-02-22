//Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu    sync.Mutex
	myMap map[string]int
}

func (c *Counter) Write(key string, wg *sync.WaitGroup) {
	c.mu.Lock()
	c.myMap[key]++
	c.mu.Unlock()
	wg.Done()
}

func (c *Counter) Read(key string) int {
	defer c.mu.Unlock()
	c.mu.Lock()
	return c.myMap[key]
}

func main() {
	key := "task7"
	wg := &sync.WaitGroup{}

	c := Counter{myMap: make(map[string]int)}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go c.Write(key, wg)
	}
	wg.Wait()
	fmt.Println(c.Read(key))
}
