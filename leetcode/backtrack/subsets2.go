/**
 * 90. Subsets II

 Given an integer array nums that may contain duplicates, return all possible subsets (the power set).

The solution set must not contain duplicate subsets. Return the solution in any order.



Example 1:

Input: nums = [1,2,2]
Output: [[],[1],[1,2],[1,2,2],[2],[2,2]]
Example 2:

Input: nums = [0]
Output: [[],[0]]


Constraints:

1 <= nums.length <= 10
-10 <= nums[i] <= 10
*/

package backtrack

import "slices"

func subsetsWithDup(nums []int) [][]int {
    res := [][]int{}
    selected := []int{}
    slices.Sort(nums)
    backtrackSubsetsWithDup(&res, &selected, nums, 0)
    return res
}

func backtrackSubsetsWithDup(res *[][]int, selected *[]int, input []int, idx int) {
    dst := make([]int, len(*selected))
    copy(dst, *selected)
    *res = append(*res, dst)
    for i := idx; i < len(input); i++ {
        if i != idx && input[i] == input[i - 1] {
            continue
        }
        *selected = append(*selected, input[i])
        backtrackSubsetsWithDup(res, selected, input, i + 1)
        *selected = (*selected)[:len(*selected) - 1]
    }
}

func subsetsWithDupUsingBitmask(nums []int) [][]int {
    res := [][]int{}
    // 1. Sort the numbers so duplicates are adjacent
    slices.Sort(nums)
    
    totalSubsets := 1 << len(nums)
    for mask := 0; mask < totalSubsets; mask++ {
        isDuplicate := false
        subset := []int{}
        
        for i := 0; i < len(nums); i++ {
            // Check if the i-th bit is set in the mask
            if (mask & (1 << i)) != 0 {
                // Duplicate elimination condition:
                // If this element is the same as the previous one,
                // and the previous one was NOT picked (its bit is 0),
                // this entire mask will generate a duplicate subset. Skip it.
                if i > 0 && nums[i] == nums[i-1] && (mask & (1 << (i-1))) == 0 {
                    isDuplicate = true
                    break
                }
                subset = append(subset, nums[i])
            }
        }
        
        // Only append if the mask represents a unique subset configuration
        if !isDuplicate {
            res = append(res, subset)
        }
    }
    return res
}