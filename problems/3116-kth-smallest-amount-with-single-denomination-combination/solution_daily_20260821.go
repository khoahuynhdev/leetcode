package main

import (
	"math/bits"
	"sort"
)

// 3116. Kth Smallest Amount With Single Denomination Combination
//
// A value v is reachable exactly when it is a multiple of at least one coin,
// so the answer is monotone in "how many reachable values are <= x". That makes
// the problem a binary search on the answer, where the predicate is a counting
// function.
//
// count(x) = |{v <= x : v is a multiple of some coin}| is computed with
// inclusion-exclusion over the (at most 2^15) subsets of coins: a subset S
// contributes (-1)^(|S|+1) * floor(x / lcm(S)).
//
// Two things keep this cheap:
//   - coins that are multiples of another coin are redundant and are dropped
//     first (e.g. [3,6,9] collapses to [3]);
//   - subset LCMs are precomputed once, and any LCM exceeding the search upper
//     bound is discarded since floor(x/lcm) is always 0 for those.
//
// The upper bound is minCoin*k: the multiples of the smallest coin alone already
// provide k reachable values by then.
func findKthSmallest(coins []int, k int) int64 {
	base := reduceCoins(coins)
	n := len(base)

	hi := base[0] * int64(k)
	// Sentinel for "this LCM can never contribute".
	tooBig := hi + 1

	// term is a surviving inclusion-exclusion term: sign * floor(x / lcm).
	type term struct {
		lcm  int64
		sign int64
	}
	terms := make([]term, 0, 1<<uint(n))

	lcms := make([]int64, 1<<uint(n))
	lcms[0] = 1
	for mask := 1; mask < 1<<uint(n); mask++ {
		low := bits.TrailingZeros(uint(mask))
		prev := lcms[mask^(1<<uint(low))]

		l := tooBig
		if prev <= hi {
			l = lcmCapped(prev, base[low], hi)
		}
		lcms[mask] = l

		if l <= hi {
			sign := int64(1)
			if bits.OnesCount(uint(mask))%2 == 0 {
				sign = -1
			}
			terms = append(terms, term{lcm: l, sign: sign})
		}
	}

	countUpTo := func(x int64) int64 {
		var total int64
		for _, t := range terms {
			total += t.sign * (x / t.lcm)
		}
		return total
	}

	// Smallest x in [1, hi] with countUpTo(x) >= k.
	lo := int64(1)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if countUpTo(mid) >= int64(k) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// reduceCoins returns the coins sorted ascending with every coin that is a
// multiple of a smaller kept coin removed. Such a coin adds no new reachable
// value, and dropping it shrinks the subset enumeration.
func reduceCoins(coins []int) []int64 {
	sorted := make([]int, len(coins))
	copy(sorted, coins)
	sort.Ints(sorted)

	base := make([]int64, 0, len(sorted))
	for _, c := range sorted {
		v := int64(c)
		redundant := false
		for _, b := range base {
			if v%b == 0 {
				redundant = true
				break
			}
		}
		if !redundant {
			base = append(base, v)
		}
	}
	return base
}

// lcmCapped returns lcm(a, b), or limit+1 if the true LCM exceeds limit. The
// cap also prevents the multiplication from overflowing int64.
func lcmCapped(a, b, limit int64) int64 {
	q := a / gcd(a, b)
	if q > limit/b {
		return limit + 1
	}
	l := q * b
	if l > limit {
		return limit + 1
	}
	return l
}

func gcd(a, b int64) int64 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
