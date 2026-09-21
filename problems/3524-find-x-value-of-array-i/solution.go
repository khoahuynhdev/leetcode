package main

// 3524. Find X Value of Array I
//
// Deleting a prefix and a non-overlapping suffix while keeping nums non-empty
// leaves exactly one contiguous block, so the answer for residue x is simply the
// number of non-empty subarrays whose product is congruent to x modulo k.
//
// Sweep the right endpoint j and keep dp[x] = how many subarrays ending at j have
// product x (mod k). Moving to j extends every such subarray by nums[j], sending
// residue x to (x*v)%k with v = nums[j]%k, and creates one new subarray [nums[j]]
// at residue v. Accumulate dp into the result after each step.
//
// Time: O(n*k), Space: O(k). Counts reach n*(n+1)/2 (~5e9), hence int64.
func resultArray(nums []int, k int) []int64 {
	result := make([]int64, k)
	dp := make([]int64, k)
	next := make([]int64, k)

	for _, num := range nums {
		v := num % k

		for x := range next {
			next[x] = 0
		}
		for x, count := range dp {
			if count != 0 {
				next[x*v%k] += count
			}
		}
		next[v]++

		dp, next = next, dp
		for x, count := range dp {
			result[x] += count
		}
	}

	return result
}
