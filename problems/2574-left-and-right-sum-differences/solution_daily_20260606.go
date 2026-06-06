package main

// 2574. Left and Right Sum Differences
//
// Approach: prefix sum in a single pass. Compute the total sum once, then sweep
// left to right maintaining a running `left` sum. The right sum at index i is
// derived for free as total - left - nums[i], since
// total = leftSum[i] + nums[i] + rightSum[i]. The answer is the absolute
// difference between the left and right sums.
//
// Time: O(n), Space: O(1) extra (excluding the output array).
func leftRightDifference(nums []int) []int {
	total := 0
	for _, v := range nums {
		total += v
	}

	answer := make([]int, len(nums))
	left := 0
	for i, v := range nums {
		right := total - left - v
		answer[i] = abs(left - right)
		left += v
	}
	return answer
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
