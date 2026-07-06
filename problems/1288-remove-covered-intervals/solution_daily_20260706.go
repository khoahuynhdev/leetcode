package main

import "sort"

// removeCoveredIntervals counts how many intervals remain after removing every
// interval that is fully covered by another.
//
// Approach: sort by start ascending, and on ties by end descending. After this
// sort, any earlier interval has a start <= the current start, so coverage
// depends only on the end value. Sweeping left to right while tracking the
// largest end seen (maxEnd), an interval is covered exactly when its end is
// <= maxEnd; otherwise it survives and extends maxEnd.
//
// Time: O(n log n) for the sort. Space: O(1) extra.
func removeCoveredIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] != intervals[j][0] {
			return intervals[i][0] < intervals[j][0]
		}
		// Equal starts: put the wider interval (larger end) first so it can
		// cover the narrower ones that follow.
		return intervals[i][1] > intervals[j][1]
	})

	count := 0
	maxEnd := 0
	for _, iv := range intervals {
		if iv[1] > maxEnd {
			// Not covered by anything seen so far.
			count++
			maxEnd = iv[1]
		}
	}
	return count
}
