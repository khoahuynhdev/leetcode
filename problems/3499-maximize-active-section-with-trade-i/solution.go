package main

// Approach: Run-length grouping.
//
// The number of 1s never decreases below the original count, so the base
// answer is the total number of '1's in s. A single trade can only help by
// picking a block of 1s that is surrounded by 0-blocks on both sides. When we
// zero out that middle 1-block of size m and then fill the merged 0-run, a
// region [a zeros][m ones][b zeros] becomes (a+m+b) ones. It previously held m
// ones, so the net gain is exactly a+b — the sizes of the two neighboring
// 0-blocks. Treating s as augmented with '1' at both ends means a 0-block that
// touches an end is still "surrounded" by a 1.
//
// So the answer is: totalOnes + max over every 1-block sandwiched between two
// 0-blocks of (leftZeros + rightZeros). If no such sandwich exists, no trade
// helps and the answer is just totalOnes.
//
// Time: O(n), Space: O(1).
func maxActiveSectionsAfterTrade(s string) int {
	n := len(s)

	totalOnes := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			totalOnes++
		}
	}

	best := 0
	i := 0
	prevZeros := -1 // length of the zero-block immediately before the current run; -1 means none
	for i < n {
		j := i
		for j < n && s[j] == s[i] {
			j++
		}
		runLen := j - i
		if s[i] == '0' {
			// If this zero-block was preceded by (a 1-block preceded by) a
			// zero-block, we have a valid sandwich; capture the gain.
			if prevZeros >= 0 && prevZeros+runLen > best {
				best = prevZeros + runLen
			}
			prevZeros = runLen
		}
		// A run of '1's keeps prevZeros so the next '0' run can pair with it.
		i = j
	}

	return totalOnes + best
}
