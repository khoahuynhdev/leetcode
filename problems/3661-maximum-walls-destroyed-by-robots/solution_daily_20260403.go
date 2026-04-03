package main

import "sort"

// Approach: Sort robots by position. Walls between consecutive robots can only
// be reached by those two robots (bullets are blocked by neighbors). Use DP
// where each robot chooses to fire left or right, and for each gap between
// consecutive robots, precompute reachable wall counts via binary search.
func maxWallsDestroyed(robots []int, distance []int, walls []int) int {
	n := len(robots)
	m := len(walls)
	if m == 0 {
		return 0
	}

	type robot struct{ pos, dist int }
	robs := make([]robot, n)
	for i := range robots {
		robs[i] = robot{robots[i], distance[i]}
	}
	sort.Slice(robs, func(i, j int) bool {
		return robs[i].pos < robs[j].pos
	})
	sort.Ints(walls)

	// Count walls in closed range [lo, hi] using binary search.
	countInRange := func(lo, hi int) int {
		if lo > hi {
			return 0
		}
		left := sort.SearchInts(walls, lo)
		right := sort.SearchInts(walls, hi+1)
		return right - left
	}

	// Walls at robot positions are always destroyed (distance 0).
	robotPos := make(map[int]bool, n)
	for _, r := range robs {
		robotPos[r.pos] = true
	}
	atRobot := 0
	for _, w := range walls {
		if robotPos[w] {
			atRobot++
		}
	}

	// Outer walls: only reachable by the boundary robots.
	leftOuter := countInRange(robs[0].pos-robs[0].dist, robs[0].pos-1)
	rightOuter := countInRange(robs[n-1].pos+1, robs[n-1].pos+robs[n-1].dist)

	if n == 1 {
		return maxVal(leftOuter, rightOuter) + atRobot
	}

	// dp[0] = best with current robot firing left
	// dp[1] = best with current robot firing right
	dp := [2]int{leftOuter, 0}

	for i := 1; i < n; i++ {
		gapLo := robs[i-1].pos + 1
		gapHi := robs[i].pos - 1

		// Robot i-1 fires right: reaches up to min(pos+dist, next robot - 1).
		rr := robs[i-1].pos + robs[i-1].dist
		if rr > gapHi {
			rr = gapHi
		}
		// Robot i fires left: reaches down to max(pos-dist, prev robot + 1).
		lr := robs[i].pos - robs[i].dist
		if lr < gapLo {
			lr = gapLo
		}

		wR := countInRange(gapLo, rr)         // reachable by robot i-1 (right)
		wL := countInRange(lr, gapHi)          // reachable by robot i (left)
		wBoth := countInRange(lr, rr)          // overlap
		wUnion := wR + wL - wBoth

		ndp := [2]int{
			maxVal(dp[0]+wL, dp[1]+wUnion), // robot i fires left
			maxVal(dp[0], dp[1]+wR),        // robot i fires right
		}
		dp = ndp
	}

	return maxVal(dp[0], dp[1]+rightOuter) + atRobot
}

func maxVal(a, b int) int {
	if a > b {
		return a
	}
	return b
}
