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