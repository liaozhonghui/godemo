package libin

import (
	in "github.com/liaozhonghui/godemo/geek/libin/internel"
	"os"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
