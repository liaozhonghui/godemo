package lru

import (
	"testing"
)

func TestPut(t *testing.T) {
	//
	var L = NewLruArray(2)
	L.Put(1)
	L.Put(1)
	L.PrintAllElemet()
}

func TestDelete(t *testing.T) {

}
