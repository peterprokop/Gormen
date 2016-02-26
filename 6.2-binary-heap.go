// 6.2-binary-heap
package main

import (
	"fmt"
)

func main() {

}

type BinaryHeap struct {
	elements []int
	heapSize int
	length   int
}

func NewBinaryHeap(length int) *BinaryHeap {
	p := new(BinaryHeap)
	p.elements = make([]int, length)
	p.length = length
	return p
}
