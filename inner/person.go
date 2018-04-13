package inner

import (
	"errors"
	"fmt"
)

//不能被外部调用，非inner package
type person struct {
	Name string
	Age  int
}

//能够被外部调用，但是内部的age不能被非inner package的操作
type Person struct {
	Name string
	age  int
}

type PersonService struct {
}

//开头字母大写，代表可以被外部调用，相当于public
func SayHello() {
	fmt.Println("Hello!!!")
}

//开头字母小写，代表只能内部调用，相当于private
func sayBye() {
	fmt.Println("Bye!!!")
}

//有了入参
func MakeMoney(money float64) {
	//g e f b等等都可以试下，具体Printf可以网上搜下都有那些格式化字符
	fmt.Printf("this is %g", money)
	fmt.Println()
}

//多入参，多返回
func MakeMoney2(money float64, name1, name2 string) (err error, data string) {
	return nil, name1
}

//这里的 PersonService被称为receiver type
func (PersonService) Dating1(name string) (err error, expenses float64) {
	if name == "Mary" {
		expenses = 500
	} else if name == "Rose" {
		expenses = 1000
	} else {
		expenses = 0
		err = errors.New("SSSS及闪避")
	}
	return
}

//这里稍作改变，函数内部只可以直接使用p的引用类型，但是使用的时候要注意p是否已经申明赋值，否则就尴尬了
func (p *PersonService) Dating2(name string) (error, float64) {
	expenses := 0.0
	if name == "Mary" {
		expenses = 500
	} else if name == "Rose" {
		expenses = 1000
	} else {
		return errors.New("SSSS及闪避"), 0
	}
	return nil, expenses
}
