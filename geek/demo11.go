package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}
	_, ok := interface{}(container).(map[int]string)
	if ok {
		fmt.Println("类型是map[int]string.")
	} else {
		fmt.Println("类型不是[]string.")
	}
}
