package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
)

var name string

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	// flag.Usage = func() {
	// 	fmt.Fprint(os.Stderr, "Usage of %s:\n", "question")
	// 	flag.PrintDefaults()
	// }
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
	temp()
}
func temp() {
	fmt.Print("hello.")
	path := []byte("AAAA/BBBBB")
	sep := bytes.IndexByte(path, '/')
	fmt.Printf("path %s, sep: %d", path, sep)
}
