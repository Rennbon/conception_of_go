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

func (p *PersonService) Wedding1(boy, girl string) (err error, children []*Person) {
	if boy == "Jack" && girl == "Rose" {
		children = append(children, &Person{Name: "王小二"})
	} else if boy == "Peter" && girl == "Mary" {
		children = append(children, &Person{Name: "Jack"}, &Person{Name: "Tom"})
	} else {
		err = errors.New("失散多年的亲兄妹.")
	}
	return
}

//可能有人认为这才是最有效率的，其实不然，其实引用传递传的是”标头“，wedding1的底层逻辑硬描述就是wedding2的状况，这个方法只是帮助更容易理解wedding1，因为这个东西当初也坑了我好久
//创建一个引用类型的变量时，创建的变量被称作标头，每个引用类型创建的标头值是包含一个指向底层数据结构的指针，标头专门是为了复制而设计的
func (p *PersonService) Wedding2(boy, girl string) (err error, children *[]*Person) {
	var err1 error
	var children1 []*Person
	if boy == "Jack" && girl == "Rose" {
		children1 = append(children1, &Person{Name: "王小二"})
	} else if boy == "Peter" && girl == "Mary" {
		children1 = append(children1, &Person{Name: "Jack"}, &Person{Name: "Tom"})
	} else {
		err1 = errors.New("失散多年的亲兄妹.")
	}
	return err1, &children1
}
