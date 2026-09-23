package main

// 1658. Minimum Operations to Reduce X to Zero
//
// Removals only ever come off the two ends, so the removed elements form a
// prefix plus a suffix and the kept elements form a contiguous middle
// subarray. Removing a total of x therefore means the middle sums to
// target = total - x, and minimizing the number of removals is the same as
// maximizing that middle subarray's length: the answer is n - longestLen.
//
// Every nums[i] is >= 1, so a window's sum grows when the right end extends
// and shrinks when the left end advances. That monotonicity lets a single
// sliding window find the longest subarray summing to target in O(n) time
// and O(1) space.
func minOperations(nums []int, x int) int {
	total := 0
	for _, v := range nums {
		total += v
	}

	target := total - x
	if target < 0 {
		// Even removing every element leaves x positive.
		return -1
	}
	if target == 0 {
		// The whole array must go; the matching middle subarray is empty.
		return len(nums)
	}

	longest := -1
	sum, left := 0, 0
	for right, v := range nums {
		sum += v
		// target > 0 here, so this never shrinks past an empty window.
		for sum > target {
			sum -= nums[left]
			left++
		}
		if sum == target && right-left+1 > longest {
			longest = right - left + 1
		}
	}

	if longest < 0 {
		return -1
	}
	return len(nums) - longest
}
