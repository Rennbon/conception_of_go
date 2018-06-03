package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	num   int32
	num2  int32
	num3  int32
	wg    sync.WaitGroup
	mutex sync.Mutex //互斥锁
)

func main() {
	wg.Add(2)
	go Inc()
	go Inc()
	wg.Wait()
	fmt.Println(num, ":", num2, ":", num3)
}
func Inc() {
	defer wg.Done()
	for i := 0; i < 10000; i++ {
		//num3++
		/*原子函数，单纯效率比互斥锁高，
		但是原子锁由底层硬件支持，脱离的golang机制，
		虽然效率很好，但是会因为各种硬件的问题变的？？？（重要的字符要打3遍）
		所欲对于严格要求一致的地方不适合原子函数
		*/
		atomic.AddInt32(&num, 1)

		mutex.Lock()
		{ //这对花括号只是为了标记作用域，可以不写
			num2++
		}
		mutex.Unlock()
	}
}
