package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go0()
	fmt.Println("end")
	//休眠阻塞1秒钟
	time.Sleep(1 * time.Second)
}
func go0() {
	go func() {
		fmt.Println("say hello")
	}()
}
