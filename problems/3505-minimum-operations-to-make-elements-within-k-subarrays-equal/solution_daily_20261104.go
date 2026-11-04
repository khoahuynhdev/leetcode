package main

import "sort"

// Approach:
//  1. For every window of size x, the cheapest way to make all elements equal
//     is to move them all to the window's median. We maintain the window with
//     two Fenwick trees over coordinate-compressed values (counts and sums),
//     which lets us find the median by binary lifting and compute the cost
//     sum|v - median| in O(log n) per window.
//  2. With cost[i] = cost of window starting at i, run a DP:
//     dp[j][i] = min ops to pick j non-overlapping windows inside nums[:i]
//     dp[j][i] = min(dp[j][i-1], dp[j-1][i-x] + cost[i-x]).
//
// Time: O(n log n + n*k), Space: O(n).
func minOperations(nums []int, x int, k int) int64 {
	n := len(nums)

	// Coordinate compression.
	vals := make([]int, n)
	copy(vals, nums)
	sort.Ints(vals)
	uniq := vals[:0]
	for i, v := range vals {
		if i == 0 || v != vals[i-1] {
			uniq = append(uniq, v)
		}
	}
	m := len(uniq)
	rank := func(v int) int { return sort.SearchInts(uniq, v) + 1 } // 1-indexed

	cntTree := make([]int64, m+1)
	sumTree := make([]int64, m+1)
	update := func(pos int, c, s int64) {
		for ; pos <= m; pos += pos & -pos {
			cntTree[pos] += c
			sumTree[pos] += s
		}
	}
	query := func(pos int) (int64, int64) {
		var c, s int64
		for ; pos > 0; pos -= pos & -pos {
			c += cntTree[pos]
			s += sumTree[pos]
		}
		return c, s
	}
	logM := 1
	for (1 << logM) <= m {
		logM++
	}
	// kth returns the smallest position whose prefix count >= target.
	kth := func(target int64) int {
		pos := 0
		for b := logM; b >= 0; b-- {
			next := pos + (1 << b)
			if next <= m && cntTree[next] < target {
				pos = next
				target -= cntTree[next]
			}
		}
		return pos + 1
	}

	windows := n - x + 1
	cost := make([]int64, windows)
	var total int64
	for i := 0; i < n; i++ {
		update(rank(nums[i]), 1, int64(nums[i]))
		total += int64(nums[i])
		if i >= x {
			update(rank(nums[i-x]), -1, -int64(nums[i-x]))
			total -= int64(nums[i-x])
		}
		if i >= x-1 {
			medPos := kth(int64((x + 1) / 2))
			med := int64(uniq[medPos-1])
			cntL, sumL := query(medPos)
			cntR, sumR := int64(x)-cntL, total-sumL
			cost[i-x+1] = med*cntL - sumL + sumR - med*cntR
		}
	}

	const inf = int64(1) << 62
	prev := make([]int64, n+1) // j-1 windows; j=0 costs nothing
	cur := make([]int64, n+1)
	for j := 1; j <= k; j++ {
		for i := 0; i <= n; i++ {
			cur[i] = inf
			if i > 0 {
				cur[i] = cur[i-1]
			}
			if i >= x && prev[i-x] < inf {
				if v := prev[i-x] + cost[i-x]; v < cur[i] {
					cur[i] = v
				}
			}
		}
		prev, cur = cur, prev
	}
	return prev[n]
}
