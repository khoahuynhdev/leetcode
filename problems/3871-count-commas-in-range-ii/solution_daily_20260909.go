package main

// 3871. Count Commas in Range II
//
// A number x gets its k-th comma exactly when it has at least 3k+1 digits,
// i.e. when x >= 10^(3k). So instead of summing commas per number, we sum over
// comma positions: the k-th comma contributes one unit for every x in [10^(3k), n].
//
//	answer = sum over k >= 1 of max(0, n - 10^(3k) + 1)
//
// With n <= 10^15 only the thresholds 10^3, 10^6, 10^9, 10^12 and 10^15 can
// contribute, so the loop runs at most five times.
func countCommas(n int64) int64 {
	const step = 1000

	var total int64
	// maxP is the largest threshold we could multiply further without overflowing.
	const maxP = int64(1) << 62

	for p := int64(step); p <= n; {
		total += n - p + 1
		if p > maxP/step {
			break
		}
		p *= step
	}

	return total
}
