package main

// 3753. Total Waviness of Numbers in Range II
//
// Approach: Digit DP with a prefix-difference trick.
//
// Let g(N) = total waviness summed over every integer in [0, N]. Then the
// answer for an inclusive range is g(num2) - g(num1-1).
//
// We build numbers digit by digit from the most-significant end. A digit is a
// peak/valley only when it has two real neighbors, and a peak/valley can be
// finalized the moment we know the digit immediately to its right. So we carry
// the previous two *real* digits in the DP state:
//
//	prev2 = the digit two positions back (the left neighbor of the candidate)
//	prev1 = the digit one position back  (the candidate middle digit)
//
// When we place a new digit d, if both previous real digits exist then prev1
// sits at an interior position and we can decide whether it is a peak
// (prev1 > prev2 && prev1 > d) or a valley (prev1 < prev2 && prev1 < d),
// contributing 1 to the waviness of every number that passes through this
// transition.
//
// Leading zeros are not real digits. We encode "no real digit placed yet" as
// the sentinel value 10, which also makes the first and last digits naturally
// ineligible to be peaks/valleys (the first digit is never the candidate with a
// left neighbor, and the last digit never gets a right neighbor).

func totalWaviness(num1 int64, num2 int64) int64 {
	return countWaviness(num2) - countWaviness(num1-1)
}

// none marks "no real digit has been placed yet" (handles leading zeros).
const none = int8(10)

type stateKey struct {
	pos          int
	prev1, prev2 int8
}

type dpResult struct {
	cnt int64 // how many numbers complete from this state
	sum int64 // total waviness contributed by those numbers' remaining digits
}

// countWaviness returns g(n): the sum of waviness over all integers in [0, n].
func countWaviness(n int64) int64 {
	if n < 0 {
		return 0
	}

	// Decompose n into its decimal digits, most-significant first.
	var digits []int8
	if n == 0 {
		digits = []int8{0}
	} else {
		for n > 0 {
			digits = append(digits, int8(n%10))
			n /= 10
		}
		for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
			digits[i], digits[j] = digits[j], digits[i]
		}
	}
	m := len(digits)

	// Memo only covers the "not tight" states; tight states lie on a single
	// path and are never revisited.
	memo := make(map[stateKey]dpResult)

	var dp func(pos int, tight bool, prev1, prev2 int8) dpResult
	dp = func(pos int, tight bool, prev1, prev2 int8) dpResult {
		if pos == m {
			return dpResult{cnt: 1, sum: 0}
		}
		if !tight {
			if v, ok := memo[stateKey{pos, prev1, prev2}]; ok {
				return v
			}
		}

		limit := int8(9)
		if tight {
			limit = digits[pos]
		}

		var res dpResult
		for d := int8(0); d <= limit; d++ {
			newTight := tight && d == limit

			var np1, np2 int8
			var contrib int64
			if prev1 == none && d == 0 {
				// Still inside the leading-zero prefix; no real digit yet.
				np1, np2 = none, none
			} else {
				// d is a real digit. If two real digits precede it, prev1 is an
				// interior digit and may be a peak or a valley.
				if prev2 != none {
					if (prev1 > prev2 && prev1 > d) || (prev1 < prev2 && prev1 < d) {
						contrib = 1
					}
				}
				np2 = prev1
				np1 = d
			}

			child := dp(pos+1, newTight, np1, np2)
			res.cnt += child.cnt
			res.sum += child.sum + contrib*child.cnt
		}

		if !tight {
			memo[stateKey{pos, prev1, prev2}] = res
		}
		return res
	}

	return dp(0, true, none, none).sum
}
