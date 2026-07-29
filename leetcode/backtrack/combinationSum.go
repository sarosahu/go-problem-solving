/**
 * 39. Combination Sum
 * Given an array of distinct integers candidates and a target integer target,
 * return a list of all unique combinations of candidates where the chosen numbers
 * sum to target. You may return the combinations in any order.

The same number may be chosen from candidates an unlimited number of times.
Two combinations are unique if the frequency of at least one of the chosen numbers is different.

The test cases are generated such that the number of unique combinations that sum up to target
is less than 150 combinations for the given input.



Example 1:

Input: candidates = [2,3,6,7], target = 7
Output: [[2,2,3],[7]]
Explanation:
2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
7 is a candidate, and 7 = 7.
These are the only two combinations.
Example 2:

Input: candidates = [2,3,5], target = 8
Output: [[2,2,2,2],[2,3,3],[3,5]]
Example 3:

Input: candidates = [2], target = 1
Output: []
**/

package backtrack

import "slices"

func combinationSum(candidates []int, target int) [][]int {
    slices.Sort(candidates)
    res := [][]int{}
    comb := []int{}
    backtrackCombinationSum(candidates, target, &res, comb, 0)
    return res
}

func backtrackCombinationSum(candidates []int, remaining int,
               res *[][]int, comb []int, idx int) {
    if remaining == 0 {
        dst := make([]int, len(comb))
        copy(dst, comb)
        *res = append(*res, dst)
        return
    }
    for i := idx; i < len(candidates) && remaining-candidates[i] >= 0; i++ {
        comb = append(comb, candidates[i])
        backtrackCombinationSum(candidates, remaining-candidates[i], res, comb, i)
        comb = comb[:len(comb) - 1]
    }
}