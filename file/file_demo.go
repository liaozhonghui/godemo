package file

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

// WriteFile 覆盖写入文件
func WriteFile(fileName, data string) {
	err := ioutil.WriteFile(fileName, []byte(data), os.ModePerm)
	check(err)
}

// AppendToFile 追加写入文件
func AppendToFile(fileName, data string) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	defer file.Close()
	check(err)
	file.Write([]byte(data))
}

// CreateFile 创建文件
func CreateFile(fileName string) {
	file, err := os.Create(fileName)
	defer file.Close()
	check(err)
}

// MkOneDir 创建单个文件夹
func MkOneDir(dir string) {
	err := os.Mkdir(dir, os.ModePerm)
	check(err)
	os.RemoveAll(dir)
}

// MkAllDir 创建多层文件夹
func MkAllDir(dirs string) {
	if !IsExist(dirs) {
		err := os.MkdirAll(dirs, os.ModePerm)
		check(err)
		os.RemoveAll(strings.Split(dirs, "/")[0])
	}
}

// DeleteFile 删除指定文件
func DeleteFile(fileName string) {
	err := os.Remove(fileName)
	check(err)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// IsExist 判断文件路径是否已经存在
func IsExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// PahtAbs 返回文件的绝对路径
func PahtAbs(path string) string {
	if absPath, err := filepath.Abs(path); err == nil {
		return absPath
	}
	return ""
}
