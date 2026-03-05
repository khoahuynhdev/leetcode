package main

// Approach: Count mismatches against the "starts with 0" alternating pattern.
// The mismatches against the "starts with 1" pattern is len(s) - count.
// Return the minimum of the two.
func minOperationsDaily20260305(s string) int {
	countA := 0 // mismatches against pattern "0101..."
	for i := 0; i < len(s); i++ {
		// Expected char for pattern A: '0' at even indices, '1' at odd
		expected := byte('0') + byte(i%2)
		if s[i] != expected {
			countA++
		}
	}
	countB := len(s) - countA
	if countA < countB {
		return countA
	}
	return countB
}
