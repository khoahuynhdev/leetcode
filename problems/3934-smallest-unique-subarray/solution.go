package main

// Approach: binary search on the answer + double rolling hash.
//
// Monotonicity: if some subarray of length L occurs exactly once, then any
// extension of it to length L+1 also occurs exactly once (every occurrence of
// the longer subarray contains an occurrence of the shorter one). The whole
// array (length n) is always unique, so an answer always exists.
//
// For a candidate length L, compute the hash of every window of length L with
// two independent polynomial hashes (mod 2^64 and mod 1e9+7), count them in a
// map, and check whether any hash appears exactly once.
//
// Time: O(n log n) expected. Space: O(n).

const (
	base1 = 1_000_003
	base2 = 911_382_323
	mod2  = 1_000_000_007
)

type hashKey struct {
	h1 uint64
	h2 uint64
}

func smallestUniqueSubarray(nums []int) int {
	n := len(nums)

	// Prefix hashes and base powers.
	pre1 := make([]uint64, n+1)
	pre2 := make([]uint64, n+1)
	pow1 := make([]uint64, n+1)
	pow2 := make([]uint64, n+1)
	pow1[0], pow2[0] = 1, 1
	for i, v := range nums {
		x := uint64(v)
		pre1[i+1] = pre1[i]*base1 + x
		pre2[i+1] = (pre2[i]*base2 + x) % mod2
		pow1[i+1] = pow1[i] * base1
		pow2[i+1] = pow2[i] * base2 % mod2
	}

	window := func(i, l int) hashKey {
		h1 := pre1[i+l] - pre1[i]*pow1[l]
		h2 := (pre2[i+l] + mod2 - pre2[i]*pow2[l]%mod2) % mod2
		return hashKey{h1, h2}
	}

	hasUnique := func(l int) bool {
		count := make(map[hashKey]int, n-l+1)
		for i := 0; i+l <= n; i++ {
			count[window(i, l)]++
		}
		for _, c := range count {
			if c == 1 {
				return true
			}
		}
		return false
	}

	lo, hi := 1, n
	for lo < hi {
		mid := lo + (hi-lo)/2
		if hasUnique(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
