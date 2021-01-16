package main

import (
	"fmt"
	"runtime/debug"
)

var N = 2

func f(n int) int {
	fmt.Println("* * * * 进入f(", n, ")")
	debug.PrintStack()
	if n > 1 {
		n = n * f(n-1)
	}
	fmt.Println("* * * * 退出f(", n, ")")
	debug.PrintStack()
	return n
}

func main() {
	f(N)
}
