// Реализовать собственную функцию sleep.
package main

import "time"

/*
time.After - возвращает канал, поэтому нужно получить значение от возвращаемого канала,
и прием будет блокироваться до тех пор, пока значение не будет отправлено
*/

func mySleep(x int) {
	<-time.After(time.Second * time.Duration(x))
}

func main() {
	mySleep(3)
}
