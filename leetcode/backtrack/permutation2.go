/*
*
  - 47. Permutations II
    Given a collection of numbers, nums, that might contain duplicates, return all possible unique permutations in any order.

Example 1:

Input: nums = [1,1,2]
Output:
[[1,1,2],

	[1,2,1],
	[2,1,1]]

Example 2:

Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]

Constraints:

1 <= nums.length <= 8
-10 <= nums[i] <= 10
*/
package backtrack

import (
	"slices"
)

// In this algorithm we are mutating original array nums.
// If we don't want to change, then we can copy to a new array
// and use it. This should be clarified first prior implementation
func permuteUnique(nums []int) [][]int {
    return permuteUniqueUsingNextPermutation(nums)
}

func permuteUniqueUsingNextPermutation(arr []int) [][]int {
	res := [][]int{}
	nums := make([]int, len(arr))
	copy(nums, arr)
	
	slices.Sort(nums)

	for {
		dst := make([]int, len(nums))
		copy(dst, nums)
		res = append(res, dst)
		if !nextPermutation(nums) {
			break
		}
	}
	return res
}

func nextPermutation(nums []int) bool {
	n := len(nums)
    k := n - 2

    for k >= 0 && nums[k] >= nums[k + 1] {
        k--
    }

    // Array is completely descending
    if k <= -1 {
        slices.Reverse(nums)
        return false
    }

    for i := n - 1; i > k; i-- {
        if nums[i] > nums[k] {
            nums[i], nums[k] = nums[k], nums[i]
            break;
        }
    }
    slices.Reverse(nums[k + 1 :])
	return true
}

func permuteUniqueBacktrack(arr []int) [][]int {
    res := [][]int{}
    counter := map[int]int{}
    for _, val := range arr {
        counter[val]++
    }
    comb := []int{}
    generateHelper(&res, comb, arr, counter)

    return res
}

func generateHelper(res *[][]int, comb []int, arr []int, counter map[int]int) {
    if len(comb) == len(arr) {
        dst := make([]int, len(arr))
        copy(dst, comb)
        *res = append(*res, dst)
        return
    }

    for key, count := range counter {
        if count > 0 {
            comb = append(comb, key)
            counter[key]--
            generateHelper(res, comb, arr, counter)
            comb = comb[:len(comb) - 1]
            counter[key]++
        }
    }
}