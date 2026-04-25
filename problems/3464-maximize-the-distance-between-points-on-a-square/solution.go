package main

// LeetCode 3464 — Maximize the Distance Between Points on a Square.
//
// Each point on the boundary maps to a single perimeter coordinate t in
// [0, 4*side). The problem becomes "place k points on a ring of length
// 4*side so the smallest cyclic gap is as large as possible."
//
// We binary search on the answer d. For d <= side (which always holds when
// k >= 4, since k*d <= 4*side by the pigeonhole), perimeter gap >= d
// implies Manhattan distance >= d:
//   - same/adjacent edges: Manhattan distance equals the perimeter gap.
//   - opposite edges: Manhattan distance >= side >= d for free.
// So the feasibility check is "sort by t, fix a start, greedily jump by
// at least d, see if we fit k picks within one lap." We double the sorted
// array and binary-search the next pick to handle the cyclic wrap.

import "sort"

func maxDistance(side int, points [][]int, k int) int {
	n := len(points)
	s := int64(side)
	P := 4 * s

	perim := make([]int64, n)
	for i, p := range points {
		x, y := int64(p[0]), int64(p[1])
		switch {
		case y == 0:
			perim[i] = x
		case x == s:
			perim[i] = s + y
		case y == s:
			perim[i] = 3*s - x
		default: // x == 0
			perim[i] = 4*s - y
		}
	}
	sort.Slice(perim, func(i, j int) bool { return perim[i] < perim[j] })

	perim2 := make([]int64, 2*n)
	for i := 0; i < n; i++ {
		perim2[i] = perim[i]
		perim2[i+n] = perim[i] + P
	}

	canPlace := func(d int64) bool {
		for i := 0; i < n; i++ {
			first := perim[i]
			limit := first + P - d
			last := first
			ok := true
			for j := 1; j < k; j++ {
				target := last + d
				lo, hi := i+1, 2*n
				for lo < hi {
					mid := (lo + hi) >> 1
					if perim2[mid] >= target {
						hi = mid
					} else {
						lo = mid + 1
					}
				}
				if lo >= 2*n || perim2[lo] > limit {
					ok = false
					break
				}
				last = perim2[lo]
			}
			if ok {
				return true
			}
		}
		return false
	}

	lo, hi := int64(0), P/int64(k)
	for lo < hi {
		mid := (lo + hi + 1) >> 1
		if canPlace(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return int(lo)
}
