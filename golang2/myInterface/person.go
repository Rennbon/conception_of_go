package myInterface

import (
	"fmt"
)

type Distincter interface {
	Distinct() error
}

type Person struct {
	Name   string
	Age    int
	PCount int
	Speed  int
}
type Chinese struct {
	Person
	Id string
}

//实现了io.Writer结构
func (o *Person) Write(p []byte) (int, error) {
	o.PCount++
	o.Name = string(p) + "-" + o.Name
	return len(p), nil
}
func (p *Person) Notify(msg string) (int, error) {
	fmt.Println(p.Name, ":", msg)
	return 0, nil
}
