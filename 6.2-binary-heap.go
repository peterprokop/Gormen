// 6.2-binary-heap
package main

import (
	"fmt"
)

func main() {
	heap := NewBinaryHeap(100)

	fmt.Println("heap.rightChild(5): ", heap.rightChild(5))

	wrongHeap := NewBinaryHeap(10)
	wrongHeap.heapSize = 10
	wrongHeap.elements = []int{0, -99, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	wrongHeap.maxHeapify(1)
	fmt.Println("wrongHeap.maxHeapify: ", wrongHeap)

	unsortedSlice := []int{-99, 10, 4, 84, 13, 55, -2, 1, 0}
	maxHeap := BuildMaxHeap(unsortedSlice)
	fmt.Println("maxHeap: ", maxHeap)

	fmt.Println("heapSort:", heapSort(unsortedSlice))
}

type BinaryHeap struct {
	elements []int // Zeroth element is ignored
	heapSize int
	length   int
}

func NewBinaryHeap(length int) *BinaryHeap {
	p := new(BinaryHeap)
	p.elements = make([]int, length)
	p.length = length
	return p
}

func BuildMaxHeap(A []int) *BinaryHeap {
	n := len(A)

	p := new(BinaryHeap)
	p.elements = make([]int, n+1)

	for i := 1; i <= n; i++ {
		p.elements[i] = A[i-1]
	}

	p.length = n
	p.heapSize = n

	for i := n / 2; i >= 1; i-- {
		p.maxHeapify(i)
	}

	return p
}

func (heap BinaryHeap) parent(i int) int {
	return i / 2
}

func (heap BinaryHeap) leftChild(i int) int {
	return i * 2
}

func (heap BinaryHeap) rightChild(i int) int {
	return i*2 + 1
}

func (heap BinaryHeap) maxHeapify(i int) {
	l := heap.leftChild(i)
	r := heap.rightChild(i)
	var largest int
	if l <= heap.heapSize && heap.elements[l] > heap.elements[i] {
		largest = l
	} else {
		largest = i
	}

	if r <= heap.heapSize && heap.elements[r] > heap.elements[largest] {
		largest = r
	}

	if largest != i {
		tmp := heap.elements[i]
		heap.elements[i] = heap.elements[largest]
		heap.elements[largest] = tmp

		heap.maxHeapify(largest)
	}
}

func heapSort(A []int) []int {
	heap := BuildMaxHeap(A)

	for i := len(A); i >= 2; i-- {
		tmp := heap.elements[1]
		heap.elements[1] = heap.elements[i]
		heap.elements[i] = tmp
		heap.heapSize--
		heap.maxHeapify(1)
	}

	return heap.elements[1:len(A)]
}

// Exercises

// 6.1-1
// What are the minimum and maximum numbers of elements in a heap of height h?
// Answer:
// 2^h - 1 and 2^(h-1) respectively
