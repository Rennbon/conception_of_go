package inner

import (
	"fmt"
	"testing"
)

//小写的只能在package为inner下使用
func TestPerson(t *testing.T) {
	p := &person{Name: "张三", Age: 20}
	fmt.Println(p)
}
