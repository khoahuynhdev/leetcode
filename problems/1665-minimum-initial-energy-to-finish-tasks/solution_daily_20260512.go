package main

// 1665. Minimum Initial Energy to Finish Tasks
//
// Approach: greedy with a custom sort.
//
// For a fixed processing order t_1..t_n, the energy spent before task k is the
// sum of the earlier tasks' "actual" costs, S_k. To legally start task k we
// need starting energy E >= minimum_k + S_k, so the minimum feasible E for that
// order is max_k (minimum_k + S_k). The total of all "actual" costs is constant
// regardless of order, so reordering only changes which minimum_k is paired with
// which prefix sum. An adjacent-swap exchange argument shows the maximum is
// minimized by processing tasks in DESCENDING order of (minimum - actual):
// finish the tasks that leave the least slack relative to their cost while
// plenty of energy is still in reserve.
//
// Then a single sweep keeps a running sum of "actual" costs and tracks the
// largest (runningSum + minimum) seen, which is the answer.
//
// Time:  O(n log n) for the sort (the sweep is O(n)).
// Space: O(log n) for the in-place sort's recursion; O(1) otherwise.

import "sort"

func minimumEffort(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1]-tasks[i][0] > tasks[j][1]-tasks[j][0]
	})

	ans := 0
	spent := 0
	for _, t := range tasks {
		actual, minimum := t[0], t[1]
		if need := spent + minimum; need > ans {
			ans = need
		}
		spent += actual
	}
	return ans
}
