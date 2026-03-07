package main

// Sliding window on doubled string.
// Double s to represent all rotations as contiguous substrings of length n.
// Compare each window against two alternating targets ("0101..." and "1010...")
// and track the minimum mismatches.
func minFlips(s string) int {
	n := len(s)
	ss := s + s
	diffA, diffB := 0, 0

	// Target A: position i expects '0' if i is even, '1' if i is odd
	// Target B: position i expects '1' if i is even, '0' if i is odd
	// A mismatch with target A at position i means match with target B, and vice versa.

	res := n // worst case: flip every character

	for i := 0; i < len(ss); i++ {
		// Expected char for target A at position i
		var expectA byte = '0'
		if i%2 == 1 {
			expectA = '1'
		}

		if ss[i] != expectA {
			diffA++
		} else {
			diffB++
		}

		// Remove the element leaving the window
		if i >= n {
			j := i - n
			var expectAj byte = '0'
			if j%2 == 1 {
				expectAj = '1'
			}
			if ss[j] != expectAj {
				diffA--
			} else {
				diffB--
			}
		}

		// Once we have a full window, check the result
		if i >= n-1 {
			if diffA < res {
				res = diffA
			}
			if diffB < res {
				res = diffB
			}
		}
	}

	return res
}
