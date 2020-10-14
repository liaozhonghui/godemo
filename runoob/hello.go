package runoob

import (
	"fmt"
)

// Hello Hello, World程序
func Hello() {
	const (
		a = iota
		b = "ha"
		c
		d
	)
	fmt.Println("Hello, World!", a, b, c, d)
}
