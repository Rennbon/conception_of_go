package myInterface

import (
	"fmt"
)

type Saver interface {
	Save() error
}

type Mongo struct {
}
type Mysql struct {
}

func (*Mongo) Save() error {
	fmt.Println("使用了mongodb存储")
	return nil
}
func (*Mysql) Save() error {
	fmt.Println("使用了mysql存储")
	return nil
}
func (*Mongo) Notify(msg string) (int, error) {
	fmt.Println("mongo save", msg)
	return len(msg), nil
}
func (*Mysql) Notify(msg string) (int, error) {
	fmt.Println("mysql save", msg)
	return len(msg), nil
}
