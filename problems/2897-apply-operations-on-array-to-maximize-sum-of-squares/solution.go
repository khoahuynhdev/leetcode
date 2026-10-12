package main

// Approach: Bit counting + greedy.
// The operation (a, b) -> (a AND b, a OR b) never changes how many elements
// have each bit set; it only moves set bits from one element to another.
// Since squares are convex, the sum of squares is maximized by concentrating
// bits into as few elements as possible. So count, for each bit position, how
// many elements have it set, then build k numbers greedily: each new number
// takes every bit that still has a remaining count.
// Time: O(n * B + k * B), Space: O(B), where B = 30 bits.
func maxSum(nums []int, k int) int {
	const mod = 1_000_000_007
	const bits = 31

	var cnt [bits]int
	for _, x := range nums {
		for b := 0; b < bits; b++ {
			if x>>b&1 == 1 {
				cnt[b]++
			}
		}
	}

	ans := 0
	for i := 0; i < k; i++ {
		x := 0
		for b := 0; b < bits; b++ {
			if cnt[b] > 0 {
				x |= 1 << b
				cnt[b]--
			}
		}
		x %= mod
		ans = (ans + x*x) % mod
	}
	return ans
}
