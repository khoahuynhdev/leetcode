package main

import "sort"

// countServers answers each query offline using a sliding window.
// Logs are sorted by time and queries are handled in increasing time order.
// Two forward-only pointers keep the window [q-x, q] over the sorted logs.
// cnt[server] holds how many logs each server has in the window, and active
// is the number of servers with cnt > 0. Each answer is n - active.
// Time: O(L log L + Q log Q + n), Space: O(n + Q).
func countServers(n int, logs [][]int, x int, queries []int) []int {
	sorted := make([][]int, len(logs))
	copy(sorted, logs)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i][1] < sorted[j][1] })

	order := make([]int, len(queries))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool { return queries[order[i]] < queries[order[j]] })

	cnt := make([]int, n+1)
	ans := make([]int, len(queries))
	active, left, right := 0, 0, 0

	for _, idx := range order {
		q := queries[idx]
		for right < len(sorted) && sorted[right][1] <= q {
			s := sorted[right][0]
			cnt[s]++
			if cnt[s] == 1 {
				active++
			}
			right++
		}
		for left < right && sorted[left][1] < q-x {
			s := sorted[left][0]
			cnt[s]--
			if cnt[s] == 0 {
				active--
			}
			left++
		}
		ans[idx] = n - active
	}
	return ans
}
