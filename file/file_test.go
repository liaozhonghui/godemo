package file

import (
	"testing"
)

const fileName = "demo.txt"

func TestCreateFile(t *testing.T) {
	CreateFile(fileName)
}

func TestWriteFile(t *testing.T) {
	WriteFile(fileName, "Hello, World")
}

func TestAppendToFile(t *testing.T) {
	AppendToFile(fileName, "First Line.")
}

func TestReadFile(t *testing.T) {
	ReadFile(fileName)
}

func TestMkdirFile(t *testing.T) {
	MkOneDir("demo")
	MkAllDir("test/user/one")
}

func TestReadAllDir(t *testing.T) {
	ReadAllDir(".")
}

func TestPathAbs(t *testing.T) {
	t.Log(PahtAbs("file"))
}
