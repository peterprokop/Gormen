package main

import "fmt"

func insertionSort(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		// Insert A[j] into the sorted sequence A[1 .. j - 1].
		i := j - 1
		for i >= 0 && A[i] > key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}

	return A
}

func main() {
	A := [7]int{5, 6, 1, 7, 4, 9, 0}

	fmt.Println("Original:", A)
	fmt.Println("Sorted:", insertionSort(A[:]))
	fmt.Println("Sorted non-increasing:", insertionSortNonIncreasing(A[:]))

	fmt.Println("Linear search 9:", linearSearch(A[:], 9))
	fmt.Println("Linear search 11:", linearSearch(A[:], 11))
}

// 2.1-2
func insertionSortNonIncreasing(A []int) []int {
	for j := 1; j < len(A); j++ {
		key := A[j]
		// Insert A[j] into the sorted sequence A[1 .. j - 1].
		i := j - 1
		for i >= 0 && A[i] < key {
			A[i+1] = A[i]
			i--
		}
		A[i+1] = key
	}

	return A
}

// 2.1-3
func linearSearch(A []int, v int) int {
	for i := 0; i < len(A); i++ {
		if A[i] == v {
			return i
		}
	}

	return -1
}

// 2.1-4
