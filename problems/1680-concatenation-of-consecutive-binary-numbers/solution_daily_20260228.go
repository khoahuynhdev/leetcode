package main

// Concatenation of Consecutive Binary Numbers
//
// For each number i from 1 to n, we "concatenate" its binary representation
// onto the running result by left-shifting the result by the bit length of i
// and adding i. The bit length only increases at powers of 2, detected via
// the bit trick i & (i-1) == 0. Modular arithmetic is applied at each step.
func concatenatedBinary(n int) int {
	const mod = 1_000_000_007
	result := 0
	bitLen := 0

	for i := 1; i <= n; i++ {
		// bit length increases at every power of 2
		if i&(i-1) == 0 {
			bitLen++
		}
		result = ((result << bitLen) + i) % mod
	}

	return result
}
