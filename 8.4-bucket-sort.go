// 8.4-bucket-sort
package main

import (
	"fmt"
)

func main() {
	A := []float32{0.9, 0.11, 0.8, 0.22, 0.7, 0.33, 0.6, 0.44,
		0.5, 0.55, 0.4, 0.66, 0.3, 0.77, 0.2, 0.88, 0.1, 0.99, 0.96}

	fmt.Println("A:", A)
	fmt.Println("A bucket-sorted:", bucketSort(A))

	L := arrayToList(A)

	fmt.Println("arrayToList(A)")
	printList(L)

	fmt.Println("insertionSortList(arrayToList(A))")
	printList(insertionSortList(L))

}

type BucketSortNode struct {
	value float32
	next  *BucketSortNode
}

func arrayToList(array []float32) *BucketSortNode {
	var firstNode *BucketSortNode = nil
	var currentNode *BucketSortNode = nil

	for i := 0; i < len(array); i++ {
		newNode := &BucketSortNode{0, nil}
		newNode.value = array[i]

		if i == 0 {
			firstNode = newNode
		} else {
			currentNode.next = newNode
		}

		currentNode = newNode
	}

	return firstNode
}

func printList(head *BucketSortNode) {
	if head != nil {
		fmt.Print(head.value, " -> ")
		printList(head.next)
	} else {
		fmt.Println(nil)
	}
}

func insertionSortList(head *BucketSortNode) *BucketSortNode {
	if head == nil {
		return nil
	}

	sortedTail := head

	for unsortedHead := head.next; unsortedHead != nil; {
		var current *BucketSortNode
		for next := head; next != nil; next = next.next {
			if next.value > unsortedHead.value {
				//fmt.Println("Non-tail")
				// Insert 'unsortedHead' after 'place'
				tmp := unsortedHead.next

				sortedTail.next = unsortedHead.next
				unsortedHead.next = next

				if current == nil {
					head = unsortedHead
				} else {
					current.next = unsortedHead
				}

				unsortedHead = tmp
				break
			}

			if next == unsortedHead {
				//fmt.Println("Tail")
				sortedTail = unsortedHead
				unsortedHead = sortedTail.next
				break
			}

			current = next
		}
	}

	return head
}

func bucketSort(A []float32) []float32 {
	n := len(A)
	B := make([]*BucketSortNode, n)
	C := make([]float32, n)

	for i := 0; i < n; i++ {
		index := int(float32(n) * A[i])
		tmp := B[index]
		B[index] = &BucketSortNode{A[i], tmp}
	}

	for i := 0; i < n; i++ {
		B[i] = insertionSortList(B[i])
	}
	// Concatentate lists together

	j := 0
	for i := 0; i < n; i++ {

		for currentNode := B[i]; currentNode != nil; currentNode = currentNode.next {
			C[j] = currentNode.value
			j++
		}
	}

	return C
}
