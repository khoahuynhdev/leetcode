package main

import "sort"

// Weighted interval scheduling with a budget of at most 4 picks.
//
// Sort the intervals by left endpoint, so that the intervals compatible with a
// chosen one always form a suffix: after taking interval i, the first usable
// position is the first left endpoint strictly greater than r_i (touching
// endpoints count as overlapping), found with one binary search.
//
// dp[j][i] = best selection taken from sorted positions i..n-1 using at most j
// intervals, where "best" means the largest score and, on ties, the
// lexicographically smallest ascending array of original indices. Each state
// carries its own index array (at most 4 entries), so inserting the current
// index into the sub-answer's sorted array is enough - no reconstruction pass.
//
// Comparing candidates by their sorted index arrays is safe: inserting the same
// index into two sorted arrays preserves their lexicographic order unless one is
// a prefix of the other, and a prefix means a strict subset, which - since every
// weight is at least 1 - would have a strictly larger score and so could never
// be tied.
//
// Time: O(n log n). Space: O(n).

// ivl is an input interval paired with its original index.
type ivl struct {
	l, r, w int
	idx     int32
}

// sel is a chosen set of at most 4 intervals: its total score plus the original
// indices kept in ascending order.
type sel struct {
	score int64
	cnt   int8
	idx   [4]int32
}

// with returns s extended by an interval of index x and weight w, keeping idx
// sorted ascending. The caller guarantees s.cnt < 4.
func (s sel) with(x int32, w int) sel {
	out := sel{score: s.score + int64(w), cnt: s.cnt + 1}
	i := int8(0)
	for ; i < s.cnt && s.idx[i] < x; i++ {
		out.idx[i] = s.idx[i]
	}
	out.idx[i] = x
	for ; i < s.cnt; i++ {
		out.idx[i+1] = s.idx[i]
	}
	return out
}

// better reports whether a beats b: higher score wins, ties go to the
// lexicographically smaller index array (a prefix being the smaller one).
func better(a, b sel) bool {
	if a.score != b.score {
		return a.score > b.score
	}
	n := a.cnt
	if b.cnt < n {
		n = b.cnt
	}
	for i := int8(0); i < n; i++ {
		if a.idx[i] != b.idx[i] {
			return a.idx[i] < b.idx[i]
		}
	}
	return a.cnt < b.cnt
}

func maximumWeight(intervals [][]int) []int {
	const maxPicks = 4

	n := len(intervals)
	if n == 0 {
		return []int{}
	}

	items := make([]ivl, n)
	for i, in := range intervals {
		items[i] = ivl{l: in[0], r: in[1], w: in[2], idx: int32(i)}
	}
	sort.Slice(items, func(a, b int) bool { return items[a].l < items[b].l })

	lefts := make([]int, n)
	for i := range items {
		lefts[i] = items[i].l
	}

	// next[i] is the first position whose interval starts after items[i] ends.
	next := make([]int, n)
	for i := range items {
		next[i] = sort.SearchInts(lefts, items[i].r+1)
	}

	dp := make([][]sel, maxPicks+1)
	for j := range dp {
		dp[j] = make([]sel, n+1)
	}
	// dp[0][*] and dp[*][n] stay at the zero value: the empty selection.

	for j := 1; j <= maxPicks; j++ {
		for i := n - 1; i >= 0; i-- {
			best := dp[j][i+1] // skip interval i
			if take := dp[j-1][next[i]].with(items[i].idx, items[i].w); better(take, best) {
				best = take
			}
			dp[j][i] = best
		}
	}

	best := dp[maxPicks][0]
	out := make([]int, best.cnt)
	for i := range out {
		out[i] = int(best.idx[i])
	}
	return out
}
