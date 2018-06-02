package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := new(sync.WaitGroup)
	//chan设置缓冲数后，可以允许channel中的
	ch := make(chan int, 10)
	fmt.Println("len1:", len(ch))
	fmt.Println("cap:", cap(ch))
	for i := 0; i < 10; i++ {
		ch <- i

		wg.Add(1)
	}
	close(ch) //ch 关闭后，ch<-会抛panic: send on closed channel
	//ch <- 100
	time.Sleep(1e9)
	fmt.Println("len2:", len(ch))
	go func() {
		for c := range ch {
			fmt.Println(c)
			wg.Done()
		}
	}()
	wg.Wait()
}
