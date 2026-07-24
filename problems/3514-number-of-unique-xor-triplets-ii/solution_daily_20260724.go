package main

// Number of Unique XOR Triplets II
//
// Approach: Because i <= j <= k permits repeated indices and XOR is
// commutative/associative, the reachable triplet values are exactly
// { x ^ y ^ z : x, y, z in distinct(nums) }. Every value is <= 1500, so any
// XOR stays below 2048 (11 bits). We first compute the set of pairwise XORs P,
// then XOR P against the distinct values to obtain every triplet XOR, using
// fixed-size boolean arrays as bitsets.
func uniqueXorTriplets(nums []int) int {
	const limit = 2048 // next power of two above the max XOR of values <= 1500

	// Collect distinct values.
	seen := make([]bool, limit)
	distinct := make([]int, 0, len(nums))
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			distinct = append(distinct, v)
		}
	}

	// Stage 1: all pairwise XORs (includes 0 from x ^ x).
	pair := make([]bool, limit)
	for i := 0; i < len(distinct); i++ {
		for j := i; j < len(distinct); j++ {
			pair[distinct[i]^distinct[j]] = true
		}
	}

	// Stage 2: XOR each pairwise result with each distinct value -> triplet XORs.
	triplet := make([]bool, limit)
	for p := 0; p < limit; p++ {
		if !pair[p] {
			continue
		}
		for _, d := range distinct {
			triplet[p^d] = true
		}
	}

	// Count reachable triplet values.
	count := 0
	for _, ok := range triplet {
		if ok {
			count++
		}
	}
	return count
}
