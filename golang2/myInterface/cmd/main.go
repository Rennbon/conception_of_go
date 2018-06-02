package main

import (
	"fmt"

	"demo/golang2/myInterface"
)

func main() {
	var p myInterface.Person
	p.Name = "Jack"
	p.Age = 20
	fmt.Println(p)

	//因为实现了io.Write接口，所以可以调用fmt.Fprint方法
	fmt.Fprint(&p, "vip")

	fmt.Println(p)

	msg := "哈哈哈，你好"
	phone := &myInterface.Email{Name: "白菜", Id: "13700000039"}
	email := &myInterface.Phone{Name: "Rennbon", Id: "rennbon@163.com"}
	//接口调用方法1
	phone.Notify(msg)
	email.Notify(msg)
	//接口方法2次封装调用，多态
	Send(phone, msg)
	Send(email, msg)

	mongo := &myInterface.Mongo{}
	mysql := &myInterface.Mysql{}

	mongo.Save()
	mysql.Save()
	//接口方法组合调用
	NS(mongo, "我要保存啦")

	Send(&p, msg)

	//嵌套
	chinese := myInterface.Chinese{
		Id: "card:999",
		Person: myInterface.Person{
			Name: "Rennbon",
			Age:  29,
		}}

	chinese.Person.Notify(msg)
	chinese.Notify(msg)
	Send(&chinese, msg)
}

//多态
func Send(n myInterface.Notifier, msg string) {
	n.Notify(msg)
}

type NotifierSaver interface {
	myInterface.Notifier
	myInterface.Saver
}

func NS(n NotifierSaver, msg string) {
	n.Notify(msg)
	n.Save()
}
