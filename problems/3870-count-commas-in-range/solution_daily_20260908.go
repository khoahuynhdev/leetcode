package main

// 3870. Count Commas in Range
//
// A d-digit number carries (d-1)/3 commas, and (d-1)/3 is exactly the number of
// k >= 1 with d >= 3k+1 -- that is, the number of powers 10^(3k) it reaches.
// So instead of summing commas per number, sum over comma columns: the k-th
// column (thousands, millions, ...) contributes one comma for every value in
// [1, n] that is at least 10^(3k), i.e. max(0, n-10^(3k)+1) of them.
//
//	answer = sum over k >= 1 of max(0, n - 10^(3k) + 1)
//
// Runs in O(log n) with O(1) extra space.
func countCommas(n int) int {
	total := 0
	for p := 1000; p <= n; {
		total += n - p + 1
		// Stop before p*1000 would overflow; any such p already exceeds n.
		if p > n/1000 {
			break
		}
		p *= 1000
	}
	return total
}
