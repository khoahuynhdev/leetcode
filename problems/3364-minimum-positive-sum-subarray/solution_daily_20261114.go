package main

// Approach: prefix sums + enumerate window lengths.
// prefix[i] holds the sum of nums[0..i-1], so the sum of nums[i..j-1] is
// prefix[j] - prefix[i]. For every end index j and every length k in [l, r],
// compute the subarray sum in O(1) and keep the smallest one that is > 0.
// Time: O(n * (r - l + 1)), Space: O(n).
func minimumSumSubarray(nums []int, l int, r int) int {
	n := len(nums)
	prefix := make([]int, n+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + v
	}

	best := -1
	for j := l; j <= n; j++ {
		for k := l; k <= r && k <= j; k++ {
			sum := prefix[j] - prefix[j-k]
			if sum > 0 && (best == -1 || sum < best) {
				best = sum
			}
		}
	}
	return best
}
