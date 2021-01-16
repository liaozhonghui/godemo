package main

import (
	"flag"
	haha "github.com/liaozhonghui/godemo/geek/lib"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	haha.Hello(name)
}
