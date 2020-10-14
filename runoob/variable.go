package runoob

import (
	"fmt"
)

// Variable 变量类型初始化
func Variable() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v, %v, %v, %q", i, f, b, s)
}
