package main

import "math/bits"

// countPrimeSetBits counts numbers in [left, right] whose popcount is prime.
// Since right <= 10^6 < 2^20, popcount is at most 20. We encode primes up to
// 20 into a bitmask for O(1) prime checking per number.
func countPrimeSetBits(left int, right int) int {
	// Bitmask with bit i set if i is prime, for i in 0..19
	// Primes: 2, 3, 5, 7, 11, 13, 17, 19
	const primeMask = 1<<2 | 1<<3 | 1<<5 | 1<<7 | 1<<11 | 1<<13 | 1<<17 | 1<<19

	count := 0
	for i := left; i <= right; i++ {
		if primeMask>>(bits.OnesCount(uint(i)))&1 == 1 {
			count++
		}
	}
	return count
}
