package array

import (
	"container/heap"
)

// 1. Define a struct to track the elements inside our heap
// It keeps a reference to the frequency coutns map to compare element weights

type MinHeap struct {
	elements[] int
	counts map[int]int
}

// Implement container/heap interface methods
func (h MinHeap) Len() int {
	return len(h.elements)
}

func (h MinHeap) Less(i, j int) bool {
	return h.counts[h.elements[i]] < h.counts[h.elements[j]]
}

func (h MinHeap) Swap(i, j int) {
	h.elements[i], h.elements[j] = h.elements[j], h.elements[i]
}

func (h *MinHeap) Push(x any) {
	h.elements = append(h.elements, x.(int))
}

func (h *MinHeap) Pop() any {
	old := h.elements
	n := h.Len()
	item := old[n - 1]
	h.elements = old[:n-1]
	return item
}

func topKFrequentMinHeap(nums []int, k int) []int {
	// Count frequencies using a map
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	// Initialize the min-heap struct
	h := &MinHeap{
		elements: make([]int, 0, k+1),
		counts: count,
	}
	heap.Init(h)

	// Keep the heap size bounded to 'k' elements
	for num := range count {
		heap.Push(h, num)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// Extract the top K elements in reverse order (highest freq comes first)
	top := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		top[i] = heap.Pop(h).(int)
	}
	return top
}

func topKFrequentBucketSort(nums []int, k int) []int {
	// 1. Count frequencies using a map
	count := make(map[int]int)
	for _, n := range nums {
		count[n]++
	}

	// 2. Create buckets. The max possible frequency is len(nums)
	// i.e. when all the items correspond to single number
	// So we need a size of len(nums) + 1
	freq := make([][]int, len(nums) + 1)

	// Populate buckets: bucket index = frequency, value = element
	for num, cnt := range count {
		freq[cnt] = append(freq[cnt], num)
	}

	// 3. Gather the top K elements by iterating backwards
	res := make([]int, k)
	idx := 0

	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			res[idx] = n
			idx++
			if idx == k {
				return res
			}
		}
	}
	return res
}