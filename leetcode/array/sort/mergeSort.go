package sort

func MergeSort(nums []int) []int {
	if (len(nums) <= 1) {
        return nums
    }
    N := len(nums)
    tmpArr := make([]int, N)
    mergeSortHelper(nums, tmpArr, 0, N - 1)
    return nums
}

func mergeSortHelper(nums []int, tmpArr []int, left int, right int) {
    if left >= right {
        return
    }
    mid := left + (right - left) / 2
    mergeSortHelper(nums, tmpArr, left, mid);
    mergeSortHelper(nums, tmpArr, mid + 1, right)
    merge(nums, tmpArr, left, right, mid)
}

func merge(nums []int, tmpArr []int, left int, right int, mid int) {
    copyArr(nums, tmpArr, left, right)
    i := left
    j := mid + 1
    k := left
    for i <= mid && j <= right {
        if tmpArr[i] <= tmpArr[j] {
            nums[k] = tmpArr[i]
            i++
        } else {
            nums[k] = tmpArr[j]
            j++
        }
        k++
    }

    for i <= mid {
        nums[k] = tmpArr[i]
        k++
        i++
    }
    for j <= right {
        nums[k] = tmpArr[j]
        k++
        j++
    }
}

func copyArr(src []int, dest []int, left int, right int) {
    for i := left; i <= right; i++ {
        dest[i] = src[i]
    }
}