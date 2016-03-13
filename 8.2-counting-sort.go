// 8.2-counting-sort
package main

import (
	"fmt"
)

func main() {
	A := []int{4, 1, 1, 3, 2, 2, 0, 0, 0}
	sorted := countingSort(A, 4)
	fmt.Println("A sorted:", sorted)
}

func countingSort(A []int, k int) []int {
	C := make([]int, k+1)
	B := make([]int, len(A))

	for j := 0; j < len(A); j++ {
		C[A[j]]++
	}

	for i := 1; i <= k; i++ {
		C[i] += C[i-1]
	}

	for j := len(A) - 1; j >= 0; j-- {
		n := A[j]
		B[C[n]-1] = n
		C[n]--
	}

	return B
}
