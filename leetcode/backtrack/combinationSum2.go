/**
 * 40. Combination Sum II

 * Given a collection of candidate numbers (candidates) and a target number (target),
   find all unique combinations in candidates where the candidate numbers sum to target.

   Each number in candidates may only be used once in the combination.

   Note: The solution set must not contain duplicate combinations.

Example 1:

Input: candidates = [10,1,2,7,6,1,5], target = 8
Output:
[
[1,1,6],
[1,2,5],
[1,7],
[2,6]
]

Example 2:

Input: candidates = [2,5,2,1,2], target = 5
Output:
[
[1,2,2],
[5]
]

Constraints:

1 <= candidates.length <= 100
1 <= candidates[i] <= 50
1 <= target <= 30
*/

package backtrack

import "slices"

func combinationSum2(candidates []int, target int) [][]int {
    res := [][]int{}
	seq := []int{}
	slices.Sort(candidates)
	combinationSum2Helper(&res, seq, candidates, target, 0)
    return res
}

// This is backtracking function.
func combinationSum2Helper(res *[][]int, seq []int, candidates []int, remaining, currIdx int) {
	if remaining <= 0 {
		if remaining == 0 {
			// copy seq to dst
			dst := make([]int, len(seq))
			copy(dst, seq)
			*res = append(*res, dst)
		}
		return
	}

	if currIdx == len(candidates) {
		return
	}

	for i := currIdx; i < len(candidates) && remaining - candidates[i] >= 0; i++ {
		if i > currIdx && candidates[i] == candidates[i - 1] {
			continue
		}
		seq = append(seq, candidates[i])
		combinationSum2Helper(res, seq, candidates, remaining - candidates[i], i + 1)
		seq = seq[:len(seq) - 1]
	}
}