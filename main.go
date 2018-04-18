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

	//冒泡排序
	/* bubbleArr := [...]int{1, 20, 3, 15, 9, 8, 4, 17}
	bubbleSlice := bubbleArr[:] */
	bubbleSlice := []int{1, 20, 3, 15, 9, 8, 4, 17}
	fmt.Println(BubbleSort(bubbleSlice))

	//---------------进阶part1 array,slice,map,struct-------------------//
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
	//如果不需要index，可以使用“_”略过
	for index, value := range slice {
		fmt.Println("index:", index, ",value:", value)
		//fmt.Println("index:", index)
		//fmt.Println("value:", value)
	}

	//最好在slice申明的时候把能预估的cap申明好，这样能节省空间
	/* 	var slice2 []int = make([]int, 0, 10)
	   	for index, value := range slice {
	   		fmt.Println("that", &slice[index], slice[index])
	   		fmt.Println("this", &value, value)
	   		slice2 = append(slice2, value)
	   		fmt.Println(cap(slice2))
	   	}
		   fmt.Println(slice2, cap(slice2)) */

	//多次运行，map是无序的，类似redis中的set
	map1 := map[string]int{"one": 1, "two": 2}
	map1["four"] = 4
	map1["three"] = 3

	for k, v := range map1 {
		fmt.Printf("map1 key is %s,value is %d", k, v)
		fmt.Println()
	}

	//---------------进阶part2 函数-------------------//
	inner.SayHello()
	//inner.sayBye()
	inner.MakeMoney(10)

	psv := &inner.PersonService{}
	//_的妙用和if的简单介绍,注：一般C#中判断成功与否是用bool类型，go的话常见用error，当然error一般会全局定义，个人认为这样能更好的归类原因
	if err, _ := psv.Dating1("Leo"); err != nil {
		fmt.Println("dating fail ", err)
	}
	//goto的不正确用法，高中数学没考过120分以上的慎用 解开"here"和"goto here"注释后运行会造成闭环
	//here:
	_, cost := psv.Dating2("Rose")
	//fmt.Printf使用
	fmt.Printf("It cost %s %g yuan to go out with Rose.", "Leo", cost)
	fmt.Println()
	fmt.Printf("It cost %[2]s %[1]g yuan to go out with Rose.", cost, "Leo")
	fmt.Println()
	//goto here
	//---------------进阶part3 &*-------------------//
	//这个方法执行完后，细细品味会对&*有进一步的理解，其实概念上Go全是传值操作的，通俗的大家会把转指针地址理解为传引用（这个花再多时间也要去理解）
	err1, children := psv.Wedding1("Jack", "Rose")
	//err1, children := psv.Wedding("Peter", "Mary")
	if err1 == nil {
		fmt.Println(children)
		//weding1
		fmt.Println("children len is ", len(children), " cap is ", cap(children))
		for _, child := range children {
			//这里是&的，表示还是个指针，所以*后去指针的值才会没有&符号
			fmt.Println(child)
		}

		/* 	//weding2
		fmt.Println("children len is ", len(*children), " cap is ", cap(*children))
		for _, child := range *children {
			fmt.Println(child)
		} */
	} else {
		fmt.Println(err1)
	}
	//copy(children, childrenClone)

}

//如果有 init() 函数则会先执行该函数
func init() {
	fmt.Println("a1=", a1)
	fmt.Println("优先加载，一般用于某些特殊逻辑的初始化")
}

//冒泡
func BubbleSort(array []int) []int {
	flag := false
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				flag = true
			}
		}
		if !flag {
			break
		}
		flag = false
	}
	return array
}
