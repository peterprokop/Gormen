package main

import "fmt"


func main() {
	A := [7]int{1, 2, 3, 4, 1, 3, 4}

	fmt.Println("A:", A)
	fmt.Println("A Merged:", merge(A[:], 0, 3, 6))

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
	L := make([]int, n1 + 1)
	R := make([]int, n2 + 1)

	for i := 1; i < n1; i++ {
		L[i] = A[p + i - 1]
	}

	for j := 1; j < n1; j++ {
		R[j] = A[q + j]
	}

	L[n1 + 1] = MaxInt
	R[n2 + 1] = MaxInt

	for k := p; k <= r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}