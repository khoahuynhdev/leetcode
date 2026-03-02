package main

// Greedy approach: for each row position i (0-indexed), we need at least
// (n-1-i) trailing zeros. Precompute trailing zeros per row, then greedily
// find the nearest qualifying row and bubble it up using adjacent swaps.
// Time: O(n^2), Space: O(n).
func minSwaps(grid [][]int) int {
	n := len(grid)
	trailingZeros := make([]int, n)
	for i := 0; i < n; i++ {
		count := 0
		for j := n - 1; j >= 0; j-- {
			if grid[i][j] == 0 {
				count++
			} else {
				break
			}
		}
		trailingZeros[i] = count
	}

	swaps := 0
	for i := 0; i < n; i++ {
		need := n - 1 - i
		// Find the first row at or below position i with enough trailing zeros.
		found := -1
		for j := i; j < n; j++ {
			if trailingZeros[j] >= need {
				found = j
				break
			}
		}
		if found == -1 {
			return -1
		}
		// Bubble the found row up to position i.
		for j := found; j > i; j-- {
			trailingZeros[j], trailingZeros[j-1] = trailingZeros[j-1], trailingZeros[j]
			swaps++
		}
	}
	return swaps
}
