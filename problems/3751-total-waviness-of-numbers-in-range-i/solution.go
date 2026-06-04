package main

// 3751. Total Waviness of Numbers in Range I
//
// Approach: direct enumeration. The range is bounded by 10^5, so we visit
// every number in [num1, num2] and sum its waviness. The waviness of a single
// number is computed by a local scan over its interior digits: a position is a
// peak if it is strictly greater than both neighbors, a valley if strictly less
// than both. The first and last digits, and any number with fewer than 3
// digits, contribute nothing because they have no interior position.
//
// Time:  O(R * D) where R = num2-num1+1 and D = digit count (<= 6).
// Space: O(D) for a reusable digit buffer.

// totalWaviness returns the sum of waviness over all integers in [num1, num2].
func totalWaviness(num1 int, num2 int) int {
	total := 0
	for x := num1; x <= num2; x++ {
		total += waviness(x)
	}
	return total
}

// waviness counts the peaks and valleys of a single non-negative integer.
func waviness(n int) int {
	// Extract digits (least-significant first). The peak/valley test is
	// symmetric, so order does not affect the count.
	digits := make([]int, 0, 7)
	if n == 0 {
		digits = append(digits, 0)
	}
	for n > 0 {
		digits = append(digits, n%10)
		n /= 10
	}

	count := 0
	for i := 1; i+1 < len(digits); i++ {
		left, mid, right := digits[i-1], digits[i], digits[i+1]
		if mid > left && mid > right { // peak
			count++
		} else if mid < left && mid < right { // valley
			count++
		}
	}
	return count
}
