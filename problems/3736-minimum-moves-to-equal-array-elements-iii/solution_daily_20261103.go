package main

// Approach: values can only increase, so the final common value must be at
// least max(nums). Choosing exactly max(nums) minimizes the cost, and each
// element x needs (max - x) moves. Answer = n*max - sum(nums). O(n) time, O(1) space.
func minMoves(nums []int) int {
	maxVal, sum := 0, 0
	for _, x := range nums {
		sum += x
		if x > maxVal {
			maxVal = x
		}
	}
	return len(nums)*maxVal - sum
}
