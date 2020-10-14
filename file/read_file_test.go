package file

import (
	"fmt"
	"testing"
)

const n = 10000

func TestWriteFile1(t *testing.T) {
	var data string = ""
	for i := 0; i < n; i++ {
		data += "Test Line Demo:" + "\n"
	}
	var fileName = "demo.txt"
	fmt.Println(fileName)
	WriteFile(fileName, data)
}
