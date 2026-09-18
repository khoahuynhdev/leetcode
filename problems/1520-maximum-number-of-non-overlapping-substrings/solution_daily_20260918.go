package main

import "sort"

// Greedy over at most 26 candidate intervals.
//
// For every letter present in s, the smallest valid substring that starts at
// that letter's first occurrence is found by expanding [first[c], last[c]] to
// the right until it contains every occurrence of every letter inside it. If the
// expansion meets a letter whose first occurrence lies left of the start, no
// valid substring starts there and the candidate is dropped (the wider interval
// is produced by that other letter instead).
//
// Surviving candidates are laminar (nested or disjoint, never partially
// overlapping), so sorting them by right endpoint and taking every interval that
// starts after the last taken end maximizes the count while minimizing total
// length: a nested interval is always considered before the one enclosing it.
//
// Time: O(26*n). Space: O(1) auxiliary.
func maxNumOfSubstrings(s string) []string {
	const alphabet = 26

	var first, last [alphabet]int
	for i := range first {
		first[i] = -1
		last[i] = -1
	}
	for i := 0; i < len(s); i++ {
		c := s[i] - 'a'
		if first[c] == -1 {
			first[c] = i
		}
		last[c] = i
	}

	type interval struct{ start, end int }
	candidates := make([]interval, 0, alphabet)

	for c := 0; c < alphabet; c++ {
		start := first[c]
		if start == -1 {
			continue
		}

		end := last[c]
		valid := true
		for j := start; j <= end; j++ {
			cur := s[j] - 'a'
			if first[cur] < start {
				valid = false
				break
			}
			if last[cur] > end {
				end = last[cur]
			}
		}
		if valid {
			candidates = append(candidates, interval{start, end})
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].end < candidates[j].end
	})

	result := make([]string, 0, len(candidates))
	prevEnd := -1
	for _, iv := range candidates {
		if iv.start > prevEnd {
			result = append(result, s[iv.start:iv.end+1])
			prevEnd = iv.end
		}
	}

	return result
}
