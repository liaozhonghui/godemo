package main

import (
	"flag"
	libin "github.com/liaozhonghui/godemo/geek/libin"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "the greeting object.")
}

func main() {
	flag.Parse()
	libin.Hello(name)
}
