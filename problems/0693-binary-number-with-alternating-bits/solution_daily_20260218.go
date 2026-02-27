package main

// hasAlternatingBits checks whether a positive integer has alternating bits.
// Approach: XOR n with n>>1. If bits alternate, the result is all 1s.
// A number x is all 1s in binary iff x & (x+1) == 0.
func hasAlternatingBits(n int) bool {
	x := n ^ (n >> 1)
	return x&(x+1) == 0
}
