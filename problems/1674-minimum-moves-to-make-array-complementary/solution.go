package main

// 1674. Minimum Moves to Make Array Complementary
//
// Approach: difference array over all possible target sums.
//
// The array consists of n/2 independent pairs (nums[i], nums[n-1-i]); every
// pair must end up summing to the same target T, and T can only be in
// [2, 2*limit]. For a pair with sorted values lo <= hi and target T:
//   - 0 moves if lo + hi == T
//   - 1 move  if lo + 1 <= T <= hi + limit  (changing one element suffices)
//   - 2 moves otherwise
//
// Start every target with a baseline cost of 2 per pair (= n total). For each
// pair, subtract 1 over the range [lo+1, hi+limit] (one move is enough there)
// and subtract another 1 at the single point lo+hi (zero moves there). A prefix
// sum over T = 2..2*limit then gives the total cost for each T; the answer is
// the minimum. Runs in O(n + limit) time and O(limit) space.
func minMoves(nums []int, limit int) int {
	n := len(nums)
	// delta[t] holds the increment applied at target t; valid targets are
	// 2..2*limit, and indices up to 2*limit+1 are touched.
	delta := make([]int, 2*limit+2)

	for i := 0; i < n/2; i++ {
		a, b := nums[i], nums[n-1-i]
		lo, hi := a, b
		if lo > hi {
			lo, hi = hi, lo
		}
		// One move suffices on [lo+1, hi+limit].
		delta[lo+1]--
		delta[hi+limit+1]++
		// Zero moves at the exact point lo+hi (nested inside the range above).
		delta[lo+hi]--
		delta[lo+hi+1]++
	}

	ans := n // upper bound: change both elements of every pair
	cur := n // baseline: 2 moves per pair
	for t := 2; t <= 2*limit; t++ {
		cur += delta[t]
		if cur < ans {
			ans = cur
		}
	}
	return ans
}
