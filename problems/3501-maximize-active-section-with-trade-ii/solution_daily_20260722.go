package main

import "sort"

// Approach:
//
// A "trade" removes a block of 1's that is surrounded by 0's, then fills a block
// of 0's surrounded by 1's. The net effect is to merge two consecutive zero-runs
// (separated by exactly one one-run) into all 1's, gaining (size of the left
// zero-run) + (size of the right zero-run) active sections. The one-run between
// them is what we remove and then re-fill.
//
// Because the parts of s outside a query range are untouched, the answer for a
// query is simply:
//
//	answer = totalOnes(s) + bestTradeGain(range)
//
// where bestTradeGain is the maximum sum of two consecutive (clipped) zero-runs
// that lie inside the augmented substring t = '1' + s[l..r] + '1'. The augmenting
// 1's only matter for the two extreme zero-runs of the range: they may be clipped
// by the boundary yet are still "surrounded" thanks to the augmentation. Every
// interior consecutive-pair uses full run sizes, so we precompute those pair sums
// and answer range-maximum queries in O(1) with a sparse table. The two boundary
// pairs are handled directly with clipped sizes.
//
// Complexity: O((n + q) log n) time, O(n) space.
func maxActiveSectionsAfterTrade(s string, queries [][]int) []int {
	n := len(s)

	totalOnes := 0
	for i := 0; i < n; i++ {
		if s[i] == '1' {
			totalOnes++
		}
	}

	// Zero-run decomposition: zStart[i]..zEnd[i] (inclusive) are all '0'.
	var zStart, zEnd []int
	for i := 0; i < n; {
		if s[i] == '0' {
			j := i
			for j < n && s[j] == '0' {
				j++
			}
			zStart = append(zStart, i)
			zEnd = append(zEnd, j-1)
			i = j
		} else {
			i++
		}
	}
	m := len(zStart)

	size := func(i int) int { return zEnd[i] - zStart[i] + 1 }

	// Sparse table over pairSum[i] = size(i) + size(i+1) for i in [0, m-2].
	var sparse [][]int
	var logTable []int
	if m >= 2 {
		cnt := m - 1
		logTable = make([]int, cnt+1)
		for i := 2; i <= cnt; i++ {
			logTable[i] = logTable[i/2] + 1
		}
		K := logTable[cnt] + 1
		sparse = make([][]int, K)
		sparse[0] = make([]int, cnt)
		for i := 0; i < cnt; i++ {
			sparse[0][i] = size(i) + size(i+1)
		}
		for k := 1; k < K; k++ {
			width := 1 << k
			sparse[k] = make([]int, cnt-width+1)
			for i := 0; i+width <= cnt; i++ {
				a := sparse[k-1][i]
				b := sparse[k-1][i+(1<<(k-1))]
				if a > b {
					sparse[k][i] = a
				} else {
					sparse[k][i] = b
				}
			}
		}
	}

	// rangeMaxPair returns the max pairSum over indices [lo, hi] (inclusive),
	// or 0 when the range is empty.
	rangeMaxPair := func(lo, hi int) int {
		if lo > hi {
			return 0
		}
		k := logTable[hi-lo+1]
		a := sparse[k][lo]
		b := sparse[k][hi-(1<<k)+1]
		if a > b {
			return a
		}
		return b
	}

	// Length of zero-run i clipped to [l, r].
	clip := func(i, l, r int) int {
		e := zEnd[i]
		if e > r {
			e = r
		}
		st := zStart[i]
		if st < l {
			st = l
		}
		return e - st + 1
	}

	ans := make([]int, len(queries))
	for qi, q := range queries {
		l, r := q[0], q[1]
		gain := 0
		if m >= 2 {
			// First zero-run whose end >= l, last zero-run whose start <= r.
			a := sort.Search(m, func(i int) bool { return zEnd[i] >= l })
			b := sort.Search(m, func(i int) bool { return zStart[i] > r }) - 1
			if a <= b && a >= 0 && b > a {
				if b == a+1 {
					// Exactly two zero-runs: both are boundary/clipped.
					gain = clip(a, l, r) + clip(b, l, r)
				} else {
					// Left boundary pair, right boundary pair, and interior pairs.
					gain = clip(a, l, r) + size(a+1)
					if g := size(b-1) + clip(b, l, r); g > gain {
						gain = g
					}
					if g := rangeMaxPair(a+1, b-2); g > gain {
						gain = g
					}
				}
			}
		}
		ans[qi] = totalOnes + gain
	}
	return ans
}
