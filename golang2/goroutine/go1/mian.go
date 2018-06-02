package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	go1()
	fmt.Println("end")
	//休眠阻塞1秒钟，9个0的纳秒
	time.Sleep(1e9)
}

//goroutine的简单使用
func go1() {
	for i := 0; i < 10; i++ {
		if i == 0 || i == 9 {
			fmt.Println("指针:", &i, ",值", i)
		}
		go func(num int) {
			//以下2个实例中num才是正确的打开方式
			fmt.Println(num)
			//下面这行中直接用i是错误的示例，当前i只是一个复制实际当前遍历值的一个中转指针
			//fmt.Println(i, &i)
		}(i) //这里的i是将当前i指向的地址传到goroutine中，具体实现原理可以去看一下源码解释，会有更深刻的了解
	}
}
