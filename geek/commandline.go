package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	//
	flag.StringVar(&name, "name", "everyone", "this greeting object.")
}

// Main name
func Main() {
	flag.Parse()
	fmt.Printf("get command parameter: %s", name)
}
