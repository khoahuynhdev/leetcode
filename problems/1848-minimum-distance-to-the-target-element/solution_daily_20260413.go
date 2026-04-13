package main

// Approach: single linear scan.
// For every index i where nums[i] == target, compute abs(i - start) and
// keep the running minimum. The problem guarantees target appears in nums,
// so the minimum is well-defined.
// Time: O(n), Space: O(1).

func getMinDistance(nums []int, target int, start int) int {
	best := len(nums) // strictly greater than any valid distance (max is n-1)
	for i, v := range nums {
		if v != target {
			continue
		}
		d := i - start
		if d < 0 {
			d = -d
		}
		if d < best {
			best = d
		}
	}
	return best
}
