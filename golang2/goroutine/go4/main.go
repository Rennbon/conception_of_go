package main

import (
	"fmt"
	"time"
)

//golang的实现的简易生产消费模式
func main() {

	ch := make(chan string)
	ch1 := make(chan string)

	Producer(ch, "我")
	Producer(ch1, "你")

	Consumer(ch)
	Consumer(ch1)

	//阻塞线程
	select {}

}
func Producer(ch chan<- string, tag string) {
	go func() {
		for i := 0; ; i++ {
			ch <- fmt.Sprintf("%s%d ", tag, i)
		}
	}()
}

func Consumer(ch <-chan string) {
	go func() {
		for c := range ch {
			fmt.Println(time.Now().Nanosecond(), c)
		}
	}()
}
