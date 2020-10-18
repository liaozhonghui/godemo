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
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s", "question")
		fmt.PrintDefaults()
	}
	flag.Parse()
	fmt.Printf("get parameter name: %s!\n", name)
}
