package main

import "fmt"

const MaxInt = int(^uint(0) >> 1)
const MinInt = -MaxInt - 1

func main() {
	A := []int{13, -3, -25, 20, 3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	B := []int{1, -4, 3, -4}

	fmt.Println("B: ", B)
	low, high, sum := findMaxSubarrayWrapper(B)
	fmt.Println("Max subarray in B: ", low, high, sum)

	fmt.Println("A: ", A)
	low, high, sum = findMaxSubarrayWrapper(A)
	fmt.Println("Max subarray in A: ", low, high, sum)
}

func findMaxCrossingSubarray(A []int, low, mid, high int) (int, int, int) {
	leftSum := MinInt
	sum := 0

	maxLeft := 0

	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}

	rightSum := MinInt
	sum = 0
	maxRight := 0

	for j := mid + 1; j <= high; j++ {
		sum += A[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}

	return maxLeft, maxRight, leftSum + rightSum
}

func findMaxSubarray(A []int, low, high int) (int, int, int) {
	if high == low {
		return low, high, A[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := findMaxSubarray(A, low, mid)
		rightLow, rightHigh, rightSum := findMaxSubarray(A, mid+1, high)
		crossLow, crossHigh, crossSum := findMaxCrossingSubarray(A, low, mid, high)

		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else if rightSum >= leftSum && rightSum >= crossSum {
			return rightLow, rightHigh, rightSum
		} else {
			return crossLow, crossHigh, crossSum
		}
	}
}

func findMaxSubarrayWrapper(A []int) (int, int, int) {
	return findMaxSubarray(A, 0, len(A)-1)
}
