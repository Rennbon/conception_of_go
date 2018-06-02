package main

import (
	"fmt"
	"sync"
)

func main() {
	//创建计数器
	wg := new(sync.WaitGroup)
	count := 100
	//设置计数为100
	wg.Add(count)
	for i := 0; i < count; i++ {
		wg.Done() //执行一次默认计数-1，这个是原子性的
		go func(num int) {
			fmt.Println(num)
		}(i)
	}
	//当计数器计数为0时，停止进程阻塞
	wg.Wait()
}
