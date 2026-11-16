package main

import "strconv"

// Approach: direct simulation. While the string is longer than k, split it
// into consecutive groups of size k (the last may be shorter), replace each
// group with the decimal string of its digit sum, and concatenate the results.
// Each round is a linear pass over the string, and the string shrinks quickly.
func digitSum(s string, k int) string {
	for len(s) > k {
		next := make([]byte, 0, len(s))
		for i := 0; i < len(s); i += k {
			end := min(i+k, len(s))
			sum := 0
			for j := i; j < end; j++ {
				sum += int(s[j] - '0')
			}
			next = strconv.AppendInt(next, int64(sum), 10)
		}
		s = string(next)
	}
	return s
}
