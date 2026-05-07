package main

// maxValue computes, for every index i, the maximum value reachable from i
// under the jump rules of LeetCode 3660 (Jump Game IX).
//
// For indices i < j, both jump rules collapse to the single edge condition
// nums[i] > nums[j], so the directed jump graph is actually undirected.
// Connected components turn out to be contiguous index ranges with
// non-decreasing per-component maxima. We compute them in one left-to-right
// pass using a monotonic stack of (segmentMax, segmentStart) entries: a new
// value v merges (pops) every segment whose max strictly exceeds v.
func maxValue(nums []int) []int {
	n := len(nums)
	type segment struct {
		max   int
		start int
	}
	stack := make([]segment, 0, n)
	for i, v := range nums {
		mergedMax := v
		mergedStart := i
		for len(stack) > 0 && stack[len(stack)-1].max > v {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if top.max > mergedMax {
				mergedMax = top.max
			}
			mergedStart = top.start
		}
		stack = append(stack, segment{max: mergedMax, start: mergedStart})
	}
	ans := make([]int, n)
	for k, seg := range stack {
		end := n
		if k+1 < len(stack) {
			end = stack[k+1].start
		}
		for j := seg.start; j < end; j++ {
			ans[j] = seg.max
		}
	}
	return ans
}
