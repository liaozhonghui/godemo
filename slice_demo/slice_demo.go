package slicedemo

import "fmt"

// List 列出slice长度
func List(n int) {
	var numbers = []int

}

func printSlice(x []int) {
	fmt.Println("len:%d, cap: %d, slice: %v\n", len(x), cap(x), x)
}
