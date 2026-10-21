package main

// Approach: case analysis on the first and last characters.
//
//   - Already sorted → 0.
//   - Length 2 and unsorted → -1 (the only proper substrings are single chars).
//   - If s[0] is already the global minimum, sorting s[1:] finishes in 1 op.
//     Symmetrically, if s[n-1] is the global maximum, sorting s[:n-1] works.
//   - Otherwise, if the minimum appears somewhere in s[:n-1], sort s[:n-1] to
//     bring it to the front, then sort s[1:] → 2 ops. Symmetrically for the
//     maximum appearing somewhere in s[1:].
//   - The only remaining case: the minimum appears only at the last index and
//     the maximum only at the first index. That needs 3 ops.
//
// Time: O(n), Space: O(1).
func minOperations(s string) int {
	n := len(s)

	sorted := true
	for i := 1; i < n; i++ {
		if s[i] < s[i-1] {
			sorted = false
			break
		}
	}
	if sorted {
		return 0
	}
	if n == 2 {
		return -1
	}

	minC, maxC := s[0], s[0]
	for i := 1; i < n; i++ {
		if s[i] < minC {
			minC = s[i]
		}
		if s[i] > maxC {
			maxC = s[i]
		}
	}

	if s[0] == minC || s[n-1] == maxC {
		return 1
	}

	// Does the minimum appear outside the last position, or the maximum
	// outside the first position?
	for i := 0; i < n-1; i++ {
		if s[i] == minC {
			return 2
		}
	}
	for i := 1; i < n; i++ {
		if s[i] == maxC {
			return 2
		}
	}
	return 3
}
