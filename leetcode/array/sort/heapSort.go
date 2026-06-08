package sort

func HeapSort(nums []int) []int {
    buildHeap(nums)
    n := len(nums)
    for i := n - 1; i >= 0; i-- {
        swap(nums, 0, i)
        heapifyDown(nums, 0, i - 1)
    }
    return nums
}

func buildHeap(nums []int) {
    n := len(nums)
    firstParentIdx := n / 2 - 1
    for i := firstParentIdx; i >= 0; i-- {
        heapifyDown(nums, i, n - 1)
    }
}

func heapifyDown(nums []int, currIdx int, endIdx int) {
    left := 2 * currIdx + 1
    largeIdx := currIdx
    for left <= endIdx {
        if nums[left] >= nums[largeIdx] {
            largeIdx = left
        }
        right := 2 * currIdx + 2
        if right <= endIdx && nums[right] >= nums[largeIdx] {
            largeIdx = right
        }
        if largeIdx != currIdx {
            swap(nums, currIdx, largeIdx)
            currIdx = largeIdx
            left = 2 * currIdx + 1
        } else {
            break
        }
    }
}

func swap(nums []int, left int, right int) {
    temp := nums[left]
    nums[left] = nums[right]
    nums[right] = temp
}