package main

// Approach: Greedy.
// Squaring is convex, so for a fixed digit sum the sum of squares is maximized
// by making digits as unequal as possible: use as many 9s as we can, then one
// digit holding the remainder, then zeros. Placing the larger digits first also
// yields the largest number among all optimal answers. If sum > 9*num, no good
// integer exists. Since sum >= 1, the leading digit is always non-zero.
func maxSumOfSquares(num int, sum int) string {
	if sum > 9*num {
		return ""
	}
	digits := make([]byte, num)
	for i := range digits {
		d := min(sum, 9)
		digits[i] = byte('0' + d)
		sum -= d
	}
	return string(digits)
}
