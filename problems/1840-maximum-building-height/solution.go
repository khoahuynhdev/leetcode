package main

// 1840. Maximum Building Height
//
// Approach: constraint propagation over "anchor" buildings.
//
// n can be up to 1e9, so we never materialize all heights. Only the restricted
// buildings, building 1 (forced to height 0), and building n (which can be the
// tallest) matter. We:
//  1. Build anchor points (id, maxHeight), adding (1, 0) and, if missing,
//     (n, n-1) — the largest height physically reachable at index n.
//  2. Sort by id, then propagate caps left-to-right and right-to-left so each
//     anchor holds its true maximum achievable height (adjacent buildings differ
//     by at most 1).
//  3. For each gap between consecutive anchors, heights rise with slope +1 from
//     both sides and meet at peak = (hA + hB + gap) / 2. The answer is the max
//     peak over all gaps.
//
// Time:  O(m log m) for m restrictions (sorting dominates).
// Space: O(m).

import "sort"

func maxBuilding(n int, restrictions [][]int) int {
	// Build anchor list with fresh inner slices so we never mutate the input.
	anchors := make([][2]int, 0, len(restrictions)+2)
	anchors = append(anchors, [2]int{1, 0}) // building 1 is forced to height 0
	hasN := false
	for _, r := range restrictions {
		anchors = append(anchors, [2]int{r[0], r[1]})
		if r[0] == n {
			hasN = true
		}
	}
	if !hasN {
		// n-1 is the tallest height index n could ever reach; never over-constrains.
		anchors = append(anchors, [2]int{n, n - 1})
	}

	sort.Slice(anchors, func(i, j int) bool {
		return anchors[i][0] < anchors[j][0]
	})

	m := len(anchors)

	// Forward propagation: a building is at most 1 taller per step than its left neighbor.
	for i := 1; i < m; i++ {
		diff := anchors[i][0] - anchors[i-1][0]
		if cap := anchors[i-1][1] + diff; cap < anchors[i][1] {
			anchors[i][1] = cap
		}
	}

	// Backward propagation: a later cap limits earlier buildings too.
	for i := m - 2; i >= 0; i-- {
		diff := anchors[i+1][0] - anchors[i][0]
		if cap := anchors[i+1][1] + diff; cap < anchors[i][1] {
			anchors[i][1] = cap
		}
	}

	// Meet-in-the-middle peak for every gap between consecutive anchors.
	ans := 0
	for i := 1; i < m; i++ {
		gap := anchors[i][0] - anchors[i-1][0]
		peak := (anchors[i-1][1] + anchors[i][1] + gap) / 2
		if peak > ans {
			ans = peak
		}
	}
	return ans
}
