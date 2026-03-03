package main

// findKthBit finds the kth bit in the nth binary string using recursion.
// Sn = S(n-1) + "1" + reverse(invert(S(n-1)))
// Instead of building the string, we recursively determine which section
// k falls into: left half (same as S(n-1)), middle (always '1'), or
// right half (maps to position 2^n - k in S(n-1), inverted).
// Time: O(n), Space: O(n)
func findKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	mid := 1 << (n - 1) // 2^(n-1)
	if k == mid {
		return '1'
	}
	if k < mid {
		return findKthBit(n-1, k)
	}
	// k is in the right half: position maps to (2^n - k) in S(n-1), inverted
	if findKthBit(n-1, (1<<n)-k) == '0' {
		return '1'
	}
	return '0'
}
