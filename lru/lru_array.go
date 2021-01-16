package lru

import (
	"errors"
	"fmt"
)

// LruArray lru数据结构
type LruArray struct {
	len    int
	cap    int
	arrays []int
}

// NewLruArray 创建lru结构
func NewLruArray(capacity int) *LruArray {
	return &LruArray{
		cap:    capacity,
		arrays: make([]int, capacity),
	}
}

// Put 放元素
func (L *LruArray) Put(ele int) (int, error) {
	i := L.findElement(ele)
	if i == -1 {
		insertRes := L.insertFirst(ele)
		if insertRes == -1 {
			return 0, errors.New("capacity is not enough")
		}
	} else {
		L.reset(ele, i)
	}
	return 
}

// Delete 删除元素
func (L *LruArray) Delete(ele int) {
	index := L.findElement(ele)
	if index > -1 {
		for i := index; i < L.len; i++ {
			L.arrays[i] = L.arrays[i+1]
		}
		L.len--
	}
}

// PrintAllElemet 打印所有元素
func (L *LruArray) PrintAllElemet() {
	for i := 0; i < L.len; i++ {
		fmt.Printf("%d ", L.arrays[i])
	}
}

func (L *LruArray) findElement(ele int) int {
	for i := 0; i < L.len; i++ {
		ele := L.arrays[i]
		if ele == val {
			return i
		}
	}
	return -1
}

func (L *LruArray) removeLast() {
	if !L.isEmpty() {
		L.len--
	}
}
func (L *LruArray) insertFirst(ele int) int {
	if L.len == L.cap {
		return -1
	}
	L.len++
	for i := 0; i < L.len; i++ {
		L.arrays[i+1] = L.arrays[i]
	}
	L.arrays[0] = ele
	return 0
}
func (L *LruArray) reset(ele int, index int) {
	for i := 0; i < L.index; i++ {
		L.arrays[i+1] = L.arrays[i]
	}
	L.arrays[0] = ele
}

func (L *LruArray) isEmpty() bool {
	return L.len == 0
}
