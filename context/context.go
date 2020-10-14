package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ReadFile 使用ioutil读取文件
func ReadFile(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	check(err)
	fmt.Println(string(data))
}

// ReadAllDir 读取文件夹
func ReadAllDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
func check() {

}

func main() {
	fmt.Println("test context!")
}
