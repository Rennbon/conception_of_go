package myInterface

import (
	"fmt"
)

type Notifier interface {
	Notify(msg string) (int, error)
}

type Email struct {
	Id   string
	Name string
}
type Phone struct {
	Id   string
	Name string
}

func (e *Email) Notify(msg string) (int, error) {
	fmt.Println("email:", e.Id, e.Name, ":", msg)
	return len(msg), nil
}
func (p *Phone) Notify(msg string) (int, error) {
	fmt.Println("phone", p.Id, p.Name, ":", msg)
	return len(msg), nil
}
