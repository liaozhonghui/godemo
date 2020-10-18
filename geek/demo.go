package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "someone", "this is user name.")
}
func main() {
	flag.Parse()
	fmt.Printf("get parameter name: %s!\n", name)
}
