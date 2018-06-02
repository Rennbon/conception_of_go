package main

import (
	"fmt"
)

func main() {
	lval := 0xFEDCBA98
	lval = lval << 2
	fmt.Printf("%#X", lval)
}
