package main

import (
	"container/heap"
	"fmt"
)

type MinHeap [] int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := h.Len()
	x := old[n - 1]
	*h = old[:n-1]
	return x
}

func main() {
	input := []int{12, 4, 5, 38, 2, 9}
	fmt.Println("--- Min Heap ---")
	minH := &MinHeap{}
	heap.Init(minH)

	for _, val := range input {
		heap.Push(minH, val)
	}

	// Peak the min value
	fmt.Printf("Minimum element (root) : %d\n", (*minH)[0])

	// Pop all elements to see ascending order
	fmt.Print("Popped sequence: ")
	for minH.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(minH).(int))
	}
	fmt.Println()
}