// 7.1-quicksort
package main

import (
	"fmt"
)

func main() {
	slice := []int{-343, 234, 0, 1, 7, 4, 11}
	quickSort(slice, 0, len(slice)-1)
	fmt.Println("Slice sorted:", slice)
}

func quickSort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)
	}
}

func partition(A []int, p, r int) int {
	x := A[r]
	i := p - 1

	for j := p; j <= r-1; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}

	A[i+1], A[r] = A[r], A[i+1]

	return i + 1
}
