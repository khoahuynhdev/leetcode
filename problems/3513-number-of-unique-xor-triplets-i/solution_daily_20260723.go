package main

// 3513. Number of Unique XOR Triplets I
//
// Approach: Because triplet indices satisfy i <= j <= k, indices may repeat,
// so a pair of equal picks cancels under XOR (x ^ x = 0). This makes every
// single element reachable, and the reachable set overall is exactly the set
// of XORs of one or three numbers drawn from the permutation {1, ..., n}.
//
// For n >= 3 the numbers 1 and 2 (and larger) are present, and XOR-combining
// values in [1, n] fills the whole range [0, 2^b - 1] where b is the bit
// length of n. Hence there are 2^b unique values. The tiny cases n == 1 and
// n == 2 don't fill a full power-of-two range and are handled directly.
//
// Time: O(log n) to compute the bit length of n. Space: O(1).
func uniqueXorTriplets(nums []int) int {
	n := len(nums)

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	// Bit length of n: number of bits needed to represent n.
	bits := 0
	for x := n; x > 0; x >>= 1 {
		bits++
	}
	return 1 << bits
}
