package array
/**
 * 315. Count of Smaller Numbers After Self

 * Given an integer array nums, return an integer array counts where counts[i] is the number of smaller elements to the right of nums[i].
 *
Example 1:
>>>>>>>>>
Input: nums = [5,2,6,1]
Output: [2,1,1,0]
Explanation:
To the right of 5 there are 2 smaller elements (2 and 1).
To the right of 2 there is only 1 smaller element (1).
To the right of 6 there is 1 smaller element (1).
To the right of 1 there is 0 smaller element.


Example 2:
>>>>>>>>>
Input: nums = [-1]
Output: [0]
Example 3:

Input: nums = [-1,-1]
Output: [0,0]
 

Constraints:

1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
 */

func countSmaller(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	indices := make([]int, n)
	
	// Initialize indices to [0, 1, 2, ..., n-1]
	for i := 0; i < n; i++ {
		indices[i] = i
	}
	
	// Temp slice allocated ONCE to prevent GC overhead during merges
	temp := make([]int, n)
	
	mergeSort(nums, indices, temp, 0, n-1, result)
	return result
}

func mergeSort(nums, indices, temp []int, start, end int, result []int) {
	if start >= end {
		return
	}
	
	mid := start + (end-start)/2
	mergeSort(nums, indices, temp, start, mid, result)
	mergeSort(nums, indices, temp, mid+1, end, result)
	
	merge(nums, indices, temp, start, mid, end, result)
}

func merge(nums, indices, temp []int, start, mid, end int, result []int) {
	leftPos, rightPos := start, mid+1
	tempPos := start
	rightCount := 0 // Tracks how many elements from the right side were smaller

	// Merge process
	for leftPos <= mid && rightPos <= end {
		// If right element is smaller, increment the running count
		if nums[indices[rightPos]] < nums[indices[leftPos]] {
			rightCount++
			temp[tempPos] = indices[rightPos]
			rightPos++
		} else {
			// Left element is smaller or equal; commit the accumulated rightCount
			result[indices[leftPos]] += rightCount
			temp[tempPos] = indices[leftPos]
			leftPos++
		}
		tempPos++
	}

	// Copy remaining elements from the left side
	for leftPos <= mid {
		result[indices[leftPos]] += rightCount
		temp[tempPos] = indices[leftPos]
		leftPos++
		tempPos++
	}

	// Copy remaining elements from the right side
	for rightPos <= end {
		temp[tempPos] = indices[rightPos]
		rightPos++
		tempPos++
	}

	// Reflect changes back to the main indices array
	for i := start; i <= end; i++ {
		indices[i] = temp[i]
	}
}
