package main

// 2770. Maximum Number of Jumps to Reach the Last Index
//
// Approach: bottom-up dynamic programming over array positions ("longest path
// in a DAG"). dp[i] is the maximum number of jumps with which index i can be
// reached starting from index 0, or -1 if i is unreachable. dp[0] = 0. For each
// i, scan every earlier index j: if j is reachable and the jump j -> i is legal
// (|nums[i] - nums[j]| <= target), then i can be reached in dp[j] + 1 jumps;
// keep the maximum. The answer is dp[n-1].
//
// Time:  O(n^2)
// Space: O(n)
func maximumJumps(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 1; i < n; i++ {
		dp[i] = -1
	}
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] == -1 {
				continue
			}
			diff := nums[i] - nums[j]
			if diff < 0 {
				diff = -diff
			}
			if diff <= target && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}
	return dp[n-1]
}
