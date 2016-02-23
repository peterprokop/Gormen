// 4.2-matrix-multiplication
package main

import (
	"errors"
	"fmt"
)

func main() {
	A := [][]int{
		{1, 0, 0},
		{0, 1, 0},
		{0, 0, 1}}

	B := [][]int{
		{-1, 0, -1},
		{0, -1, 0},
		{-1, 0, -1}}

	C, error := squareMatrixMultiply(A, B)
	fmt.Println("A x B, error: ", C, error)

	C, error = squareMatrixMultiply(B, B)
	fmt.Println("B x B, error: ", C, error)
}

func squareMatrixMultiplyDimensionError() ([][]int, error) {
	return make([][]int, 0), errors.New("Matrix dimensions doesn't match")
}

func squareMatrixMultiply(A, B [][]int) ([][]int, error) {

	// Check dimensions
	n := len(A)

	if n != len(B) {
		return squareMatrixMultiplyDimensionError()
	}

	for i := 0; i < n; i++ {
		if (len(A[i]) != n) || (len(B[i]) != n) {
			return squareMatrixMultiplyDimensionError()
		}
	}

	// Dimensions are OK, compute product

	C := make([][]int, n)
	for i := 0; i < n; i++ {
		C[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = 0
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}

	return C, nil
}

// TODO: SQUARE-MATRIX-MULTIPLY-RECURSIVE
// TODO: Strassen’s algorithm
