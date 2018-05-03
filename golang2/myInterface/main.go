package myInterface

import (
	"fmt"
)

func main() {
	var p Person
	p.Name = "Jack"
	p.Age = 20
	fmt.Println(p)

	//因为实现了io.Write接口，所以可以调用fmt.Fprint方法
	fmt.Fprint(&p, "vip")
	fmt.Println(p)
}
