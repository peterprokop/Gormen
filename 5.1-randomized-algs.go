package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		fmt.Println(random(1, 10))
	}
}

// Exercises

// 5.1-2
func random(a, b int) int {
	n := b - a + 1

	if n == 1 {
		return a
	}

	// User binary-search-like method if n is even
	if n%2 == 0 {
		if rand.Intn(2) == 0 {
			return random(a, a+n/2-1)
		} else {
			return random(a+n/2, b)
		}
	} else {
		sum := 0
		for i := 1; i <= n; i++ {
			sum += rand.Intn(2)
		}

		if sum == n {
			return b
		} else {
			return random(a, b-1)
		}
	}
}

// TODO: 5.1-3
// Suppose that you want to output 0 with probability 1=2 and 1 with probability 1=2.
// At your disposal is a procedure BIASED-RANDOM, that outputs either 0 or 1.
// It outputs 1 with some probability p and 0 with probability 1-p, where 0 < p < 1, but you do not know what p is.
// Give an algorithm that uses BIASED-RANDOM as a subroutine, and returns an unbiased answer, returning 0 with probability 1/2
// and 1 with probability 1/2. What is the expected running time of your algorithm as a function of p?
