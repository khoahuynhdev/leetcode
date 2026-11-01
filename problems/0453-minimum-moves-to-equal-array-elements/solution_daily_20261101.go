package main

// Approach: incrementing n-1 elements by 1 is equivalent (relative to each
// other) to decrementing a single element by 1. So the minimum number of moves
// is the number of single-element decrements needed to bring every element down
// to the array minimum: sum(nums[i] - min).
// Time: O(n), Space: O(1).
func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	minVal := nums[0]
	for _, x := range nums[1:] {
		if x < minVal {
			minVal = x
		}
	}

	var moves int64
	for _, x := range nums {
		moves += int64(x - minVal)
	}
	return int(moves)
}
