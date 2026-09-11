package main

// 3483. Unique 3-Digit Even Numbers
//
// Approach: brute-force enumeration over ordered index triples.
// Since len(digits) <= 10, we can try every ordered choice of three
// distinct positions (i, j, k) and keep the candidate when it forms a
// valid number: no leading zero (digits[i] != 0) and even (digits[k]%2 == 0).
// Indexing by position rather than by value automatically respects the
// "each copy of a digit may be used once per number" rule, while a
// seen[1000] bitmap over the resulting value collapses duplicates that
// come from repeated digits in the input.
//
// Time: O(n^3) with n <= 10, Space: O(1).
func totalNumbers(digits []int) int {
	var seen [1000]bool
	count := 0

	for i := range digits {
		if digits[i] == 0 { // no leading zeros
			continue
		}
		for j := range digits {
			if j == i {
				continue
			}
			for k := range digits {
				if k == i || k == j {
					continue
				}
				if digits[k]%2 != 0 { // must be even
					continue
				}
				value := digits[i]*100 + digits[j]*10 + digits[k]
				if !seen[value] {
					seen[value] = true
					count++
				}
			}
		}
	}

	return count
}
