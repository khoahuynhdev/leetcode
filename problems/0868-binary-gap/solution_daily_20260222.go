package main

// binaryGap finds the longest distance between any two adjacent 1's
// in the binary representation of n by iterating through each bit,
// tracking the position of the last seen 1, and computing the maximum gap.
func binaryGap(n int) int {
	maxGap := 0
	last := -1
	pos := 0

	for n > 0 {
		if n&1 == 1 {
			if last != -1 {
				gap := pos - last
				if gap > maxGap {
					maxGap = gap
				}
			}
			last = pos
		}
		n >>= 1
		pos++
	}

	return maxGap
}
