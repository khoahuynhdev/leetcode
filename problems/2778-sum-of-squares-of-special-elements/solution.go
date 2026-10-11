package main

// Approach: An element at 1-indexed position i is special when i divides n.
// Iterate over every position, and for each divisor of n add the square of
// the element at that position (nums[i-1] in 0-indexed Go). O(n) time, O(1) space.
func sumOfSquares(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			sum += nums[i-1] * nums[i-1]
		}
	}
	return sum
}
