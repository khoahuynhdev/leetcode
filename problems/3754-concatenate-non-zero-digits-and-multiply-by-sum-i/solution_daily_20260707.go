package main

import "strconv"

// Approach: Scan the decimal digits of n from most significant to least
// significant (by iterating over its string form). For each non-zero digit d,
// grow the concatenated value with x = x*10 + d and add d to sum. Zero digits
// are skipped. If there are no non-zero digits, x and sum both remain 0, so the
// answer is naturally 0. Finally return x * sum.
//
// n has at most 10 digits (0 <= n <= 10^9), so x and x*sum fit in an int.
func concatenateNonZeroDigits(n int) int {
	x, sum := 0, 0
	for _, c := range strconv.Itoa(n) {
		d := int(c - '0')
		if d != 0 {
			x = x*10 + d
			sum += d
		}
	}
	return x * sum
}
