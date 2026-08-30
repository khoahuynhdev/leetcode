package main

// Greedy, single pass.
//
// Deletions only happen at the two ends, so whatever survives is a contiguous
// middle slice nums[a : n-b] costing a+b. Locate the min and max indices and
// normalize them to lo <= hi; there are then exactly three ways to push both
// outside the surviving slice:
//
//	front only: hi + 1
//	back only:  n - lo
//	split:      (lo + 1) + (n - hi)
//
// The answer is the cheapest of the three. Runs in O(n) time, O(1) space.
func minimumDeletions(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	minIdx, maxIdx := 0, 0
	for i, v := range nums {
		if v < nums[minIdx] {
			minIdx = i
		}
		if v > nums[maxIdx] {
			maxIdx = i
		}
	}

	lo, hi := min(minIdx, maxIdx), max(minIdx, maxIdx)

	fromFront := hi + 1
	fromBack := n - lo
	split := (lo + 1) + (n - hi)

	return min(fromFront, min(fromBack, split))
}
