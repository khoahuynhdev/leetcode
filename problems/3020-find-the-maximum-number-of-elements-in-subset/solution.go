package main

// 3020. Find the Maximum Number of Elements in Subset
//
// The valid arrangement [x, x^2, x^4, ..., x^k, ..., x^4, x^2, x] is a
// palindrome built by repeated squaring. Every value below the peak appears
// twice; the peak appears once. So we count occurrences of each value, then
// for every distinct base x we walk the squaring chain upward while the
// current value occurs at least twice, and resolve the peak at the top.
//
// The value 1 is special (1^2 == 1): a chain of 1s is just an odd-length run.

func maximumLength(nums []int) int {
	freq := make(map[int]int, len(nums))
	for _, v := range nums {
		freq[v]++
	}

	ans := 1

	// Base 1 collapses to a run of 1s, which must have odd length.
	if c, ok := freq[1]; ok {
		if c%2 == 0 {
			c--
		}
		if c > ans {
			ans = c
		}
	}

	for v := range freq {
		if v == 1 {
			continue
		}
		cur := v
		length := 0
		// Each value with at least two copies contributes a left and right copy.
		for freq[cur] >= 2 {
			length += 2
			cur *= cur // safe in 64-bit: cur <= 1e9, so cur*cur <= 1e18
		}
		if freq[cur] == 1 {
			// This value serves as the unique peak.
			length++
		} else {
			// Chain ran out; the last paired value becomes a single peak.
			length--
		}
		if length > ans {
			ans = length
		}
	}

	return ans
}
