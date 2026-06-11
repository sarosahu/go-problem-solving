package array

import (
	"container/heap"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// Define the PriorityQueue type
type PriorityQueue []*ListNode

// Len returns the total no. of elements in the queue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less defines the priority.
// Use < for Min-Heap & > for Max-Heap
func (pq PriorityQueue) Less (i, j int) bool {
	return pq[i].Val < pq[j].Val
}

// Swap exchanges two elements in the slice.
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push adds an item to the end of the slice.
func (pq *PriorityQueue) Push(x any) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

// Pop removes and returns the last element of the slice
func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n - 1]
	*pq = old[0 : n - 1]
	return item
}


func MergeKLists(lists []*ListNode) *ListNode {
	// 1. Init an empty priority queue
    pq := &PriorityQueue{}
	heap.Init(pq)

	// 2. Push the head node of each non-empty list into the Min-Heap
	for _, head := range lists {
		if head != nil {
			heap.Push(pq, head)
		}
	}

	// 3. Create a dummy head node to build our result list
	dummy := &ListNode{}
	tail := dummy

	// 4. Process nodes until the heap is completely empty
	for pq.Len() > 0 {
		smallestNode := heap.Pop(pq).(*ListNode)

		// Append this smallest node to our merged list
		tail.Next = smallestNode
		tail = tail.Next

		// If this node has a next node, push that next node into the heap
		if smallestNode.Next != nil {
			heap.Push(pq, smallestNode.Next)
		}
	}

	// 5. Return the head of the newly merged sorted list
	return dummy.Next
}