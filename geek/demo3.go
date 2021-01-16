package main

import (
	"flag"
	"fmt"
	"os"
)

var name string
var cmdLine = flag.NewFlagSet("name", flag.ExitOnError)

func init() {
	cmdLine.StringVar(&name, "name", "someone", "this is user name.")
}
func main() {
	// flag.Usage = func() {
	// 	fmt.Fprintf(os.Stderr, "Usage of %s\n", "question")
	// 	flag.PrintDefaults()
	// }
	cmdLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s\n", "name")
		flag.PrintDefaults()
	}

	cmdLine.Parse(os.Args[1:])
	fmt.Printf("get parameter name: %s!\n", name)
}
