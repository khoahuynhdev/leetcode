package main

// Find the Prefix Common Array of Two Arrays (LeetCode 2657).
//
// Approach: single left-to-right pass with a frequency tally.
// Because A and B are permutations of 1..n, each value can be seen at most
// twice across the two arrays. The instant a value's tally reaches 2, it has
// appeared in both prefixes and becomes "common", so we bump a running counter.
// C[i] is simply that counter after processing index i.
//
// Time: O(n), Space: O(n).
func findThePrefixCommonArray(A []int, B []int) []int {
	n := len(A)
	seen := make([]int, n+1) // values are in 1..n; index 0 is unused
	c := make([]int, n)
	common := 0

	for i := 0; i < n; i++ {
		seen[A[i]]++
		if seen[A[i]] == 2 {
			common++
		}
		seen[B[i]]++
		if seen[B[i]] == 2 {
			common++
		}
		c[i] = common
	}

	return c
}
