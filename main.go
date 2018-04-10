/*包名*/
package main

//导包,就是引用其他包
import (
	"demo/inner"
	"fmt"
	"reflect"
)

//明确定义a1为int类型并赋值为1
var a1 int = 1

//定义a2变量，系统会按照赋值类型自动推断a2是什么类型
var a2 = "23"

var b1, b2, b3 string = "1", "2", "3"
var c1, c2, c3 = "1", 2, true

var ( // 这种因式分解关键字的写法一般用于声明全局变量
	d1 string
	d2 int
	d3 bool
)

//程序开始执行的函数
func main() {
	//程序员的世界
	fmt.Println("Hello world")
	a1 = 2
	fmt.Println("a1=", a1)

	//输出a2的推断类型
	fmt.Println(reflect.TypeOf(a2))
	//这种申明变量只能用于函数内部
	a3 := 5
	fmt.Println(a3)

	//---------------进阶part1,array,slice,map,struct-------------------//
	arr := new([10]int)
	fmt.Println("arr=", *arr)
	// 10为len,20为cap
	slice := make([]int, 10, 11)
	slice[9] = 10
	fmt.Println("slice=", slice, ",len=", len(slice), ",cap=", cap(slice))
	slice = append(slice, 20)
	//在slice 中append后超过当前cap时，cap会翻倍
	//slice = append(slice, 20)
	fmt.Println("slice=", slice, ",len=", len(slice), ",cap=", cap(slice))
	slice = slice[0:10]
	fmt.Println("slice=", slice, ",len=", len(slice), ",cap=", cap(slice))

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}
	for index, value := range slice {
		fmt.Println("index:", index, ",value:", value)
	}
	//---------------进阶part2函数-------------------//
	inner.SayHello()
	//inner.sayBye()
	inner.MakeMoney(10)

}

//如果有 init() 函数则会先执行该函数
func init() {
	fmt.Println("a1=", a1)
	fmt.Println("优先加载，一般用于某些特殊逻辑的初始化")
}
