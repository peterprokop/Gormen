// 8.2-counting-sort
package main

import (
	"fmt"
	"math"
)

func main() {
	A := []int{4, 1, 1, 3, 2, 2, 0, 0, 0}
	sorted := countingSort(A, 4)
	fmt.Println("A sorted:", sorted)

	B := []int{121, 125, 100, 190, 130, 120, 155, 160, 140, 180}
	fmt.Println("B sorted by 2 digit:", countingSortByDigit(B, 9, 2))

	fmt.Println("B radix-sorted:", radixSort(B, 9, 3))
}

/**
countingSort sorts input array *A* using counting sort
Time complexity: O(len(A) + k)
- parameter A: input array
- parameter k: values of *A* should be in *0..k*
*/
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

// 8.3-radix-sort

/**
countingSortByDigit sorts input array *A* using counting sort
Time complexity: O(len(A) + k)
- parameter A: input array
- parameter k: values of *A* should be in *0..k*
*/
func countingSortByDigit(A []int, k, d int) []int {
	C := make([]int, k+1)
	B := make([]int, len(A))
	denominator := math.Pow(float64(k+1), float64(d-1))

	for j := 0; j < len(A); j++ {
		digit := (A[j] / int(denominator)) % (k + 1)
		C[digit]++
	}

	for i := 1; i <= k; i++ {
		C[i] += C[i-1]
	}

	for j := len(A) - 1; j >= 0; j-- {
		digit := (A[j] / int(denominator)) % (k + 1)
		B[C[digit]-1] = A[j]
		C[digit]--
	}

	return B
}

func radixSort(A []int, k, d int) []int {
	B := A

	for i := 1; i <= d; i++ {
		B = countingSortByDigit(B, k, i)
	}

	return B
}
