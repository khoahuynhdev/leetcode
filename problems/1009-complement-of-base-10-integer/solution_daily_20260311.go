package main

// bitwiseComplement returns the complement of n by XORing it with a mask
// of all 1s that has the same bit-length as n.
func bitwiseComplement(n int) int {
	if n == 0 {
		return 1
	}

	// Build a mask with all 1s matching the bit-length of n.
	mask := 1
	for mask <= n {
		mask <<= 1
	}

	return n ^ (mask - 1)
}
