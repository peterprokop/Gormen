package main

import "fmt"

func main() {
	A := [7]int{1, 2, 3, 4, 1, 3, 4}

	fmt.Println("A:", A)
	fmt.Println("A Merged:", merge(A[:], 0, 3, 6))

	B := [7]int{19, 5, 3, 8, 13, 1, 6}

	fmt.Println("B:", B)
	fmt.Println("B Merge-sorted:", mergeSort(B[:]))

	C := [7]int{19, 5, 3, 8, 13, 1, 6}

	fmt.Println("C:", C)
	fmt.Println("C Bubble-sorted:", mergeSort(C[:]))

}

/*
Extreme int values in go:
const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1
*/

const MaxInt = int(^uint(0) >> 1)

func merge(A []int, p int, q int, r int) []int {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1+1)
	R := make([]int, n2+1)

	for i := 0; i < n1; i++ {
		L[i] = A[p+i]
	}

	for j := 0; j < n2; j++ {
		R[j] = A[q+j+1]
	}

	L[n1] = MaxInt
	R[n2] = MaxInt

	i := 0
	j := 0

	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}

	return A
}

func mergeSort(A []int) []int {
	return mergeSortInternal(A, 0, len(A)-1)
}

func mergeSortInternal(A []int, p int, r int) []int {
	if p < r {
		q := (p + r) / 2
		mergeSortInternal(A, p, q)
		mergeSortInternal(A, q+1, r)
		merge(A, p, q, r)
	}

	return A
}

// Problems

// 2-2
func bubbleSort(A []int) []int {
	for i := 0; i < len(A)-1; i++ {
		for j := len(A) - 1; j >= i+1; j-- {
			if A[j] < A[j-1] {
				tmp := A[j]
				A[j] = A[j-1]
				A[j-1] = tmp
			}
		}
	}

	return A
}

// 2-4

// a. List the fivei nversions of the array <2;3;8;6;1>
// Answer:
// (1,5) (2,5) (3,5) (4,5) (3,4)

// b. What array with elements from the set {1; 2; ... ; n}
// has the most inversions? How many does it have?
// Answer:
// Array sorted in reverse order. {n; ...; 2; 1}
// It's easy to show that it has (n-1) + (n-2) + ... + 1 =
// = n*(n-1)/2 inversions

// c. What is the relationship between the running time of
// insertion sort and the number of inversions in the input array? Justify your answer.
// Answer:
// Let `I` be the number of inversions and `n` number of elements in the array.
// It can be shown that running time of insertion sort is `Theta(I + n)`
