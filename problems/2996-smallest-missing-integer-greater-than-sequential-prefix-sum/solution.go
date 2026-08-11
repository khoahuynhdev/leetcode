package main

// 2996. Smallest Missing Integer Greater Than Sequential Prefix Sum
//
// Two independent passes:
//  1. Walk forward from index 0 while nums[i] == nums[i-1]+1, accumulating the
//     sum. A prefix must start at index 0, so the first violation ends the
//     longest sequential prefix for good.
//  2. Put every element in a set, then probe upward from that sum for the first
//     value not present. Values are capped at 50, so at most 51 consecutive
//     integers can be blocked and the probe is effectively O(1).
//
// Time: O(n). Space: O(n) for the set.
func missingInteger(nums []int) int {
	sum := nums[0]
	for i := 1; i < len(nums) && nums[i] == nums[i-1]+1; i++ {
		sum += nums[i]
	}

	seen := make(map[int]struct{}, len(nums))
	for _, n := range nums {
		seen[n] = struct{}{}
	}

	x := sum
	for {
		if _, ok := seen[x]; !ok {
			return x
		}
		x++
	}
}
