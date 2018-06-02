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
	defer func() {
		log.Println("done")
		if err := recover(); err != nil {
			log.Printf("run time panic: %v", err)
		}
	}()
	g()
}
