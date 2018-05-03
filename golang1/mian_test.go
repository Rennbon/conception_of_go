package main

import (
	"demo/golang1/inner"
	"fmt"
	"testing"
)

//大写的只能在package为inner下使用，但是不能对age操作
func TestPerson(t *testing.T) {
	p := &inner.Person{Name: "张三"}
	fmt.Println(p)
}
