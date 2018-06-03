package main

import "log"

func main() {
	log.Println("service start")
	//unforgivableSystemException()
	protect(unforgivableSystemException)
	log.Println("service stop")
}
func unforgivableSystemException() {
	panic("WTF！！！！")
}

func protect(g func()) {
	//执行完g()后执行
	defer func() {
		log.Println("done")
		//如果有panic的话会被recover到，相当于C#中的catch，阻止系统直接抛异常
		if err := recover(); err != nil {
			//为什么我这里没用fmt打印，而用log，你懂得，可以异常数据收集
			log.Printf("run time panic: %v", err)
		}
	}()
	//执行对应的函数
	g()
}
