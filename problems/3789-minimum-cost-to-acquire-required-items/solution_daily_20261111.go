package main

// Approach: Let k be the number of type-3 ("both") items bought. For a fixed k,
// the cheapest completion is max(0, need1-k) type-1 items and max(0, need2-k)
// type-2 items. The total cost is piecewise linear in k with breakpoints at
// min(need1, need2) and max(need1, need2), so the minimum is reached at one of
// k = 0, k = min(need1, need2), or k = max(need1, need2). Evaluate all three.
// Time O(1), space O(1). Values reach ~2e15, so int64 is used.

func minimumCost(cost1 int, cost2 int, costBoth int, need1 int, need2 int) int64 {
	c1, c2, cb := int64(cost1), int64(cost2), int64(costBoth)
	n1, n2 := int64(need1), int64(need2)

	costFor := func(k int64) int64 {
		r1 := n1 - k
		if r1 < 0 {
			r1 = 0
		}
		r2 := n2 - k
		if r2 < 0 {
			r2 = 0
		}
		return k*cb + r1*c1 + r2*c2
	}

	lo, hi := n1, n2
	if lo > hi {
		lo, hi = hi, lo
	}

	best := costFor(0)
	for _, k := range []int64{lo, hi} {
		if c := costFor(k); c < best {
			best = c
		}
	}
	return best
}
