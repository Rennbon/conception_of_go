package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	//最大pcount数就是runtime.NumCPU()的返回值

	fmt.Println(runtime.NumCPU())
	/*pcount设置为1时，gorounte会默认在一个P（处理器）上执行，设置为2以上的时候会多处理器并行
	所以当1时，大小写各自的顺序都是连续的，当大于1时，会出现大小写打乱的情况
	但是当1时，如果当前goroutine处理时间超过一定时间，调度器会从G队列中找其他可执行的对象
	*/
	pcount := 1
	runtime.GOMAXPROCS(pcount)

	print(true)
	print(false)
	time.Sleep(1e9)
}

func print(lower bool) {
	c := 'a'
	if lower {
		c = 'A'
	}
	go func() {
		for i := c; i < c+26; i++ {
			//在pcount=1时，放行下面的代码，可以体验goroutine的1P下的调度机制
			//time.Sleep(1e2)
			fmt.Printf("%c ", i)
		}
	}()
}
