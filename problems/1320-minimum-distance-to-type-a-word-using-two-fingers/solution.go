package main

// Approach: Dynamic programming over the position of the "idle" finger.
// After typing word[i], the active finger is pinned at word[i], so we only
// need to track where the other finger sits. Letter L maps to grid
// coordinate (L/6, L%6) and distance is Manhattan. We use a sentinel value
// 26 to represent "other finger not yet placed" (its first placement is
// free). Transition from dp[o] at step i to step i+1:
//   - Same (active) finger moves:   dp'[o]       = dp[o] + dist(prev, curr)
//   - Other (idle) finger moves:    dp'[prev]    = dp[o] + (o==26 ? 0 : dist(o, curr))
// Answer is the minimum over all idle-finger positions after the last letter.
// Time  O(n * 27)  ≈ O(n)
// Space O(27)     = O(1)

func minimumDistance(word string) int {
	n := len(word)
	if n <= 1 {
		return 0
	}

	const free = 26
	const inf = 1 << 30

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	dist := func(a, b int) int {
		if a == free || b == free {
			return 0
		}
		return abs(a/6-b/6) + abs(a%6-b%6)
	}

	dp := make([]int, 27)
	for i := range dp {
		dp[i] = inf
	}
	dp[free] = 0 // after typing word[0] for free, the other finger is unplaced

	next := make([]int, 27)
	for i := 1; i < n; i++ {
		prev := int(word[i-1] - 'A')
		curr := int(word[i] - 'A')
		moveSame := dist(prev, curr)

		for j := range next {
			next[j] = inf
		}

		for o := 0; o <= free; o++ {
			if dp[o] >= inf {
				continue
			}
			// Active finger moves from prev to curr; idle finger stays at o.
			if v := dp[o] + moveSame; v < next[o] {
				next[o] = v
			}
			// Idle finger at o moves to curr; the previously-active finger
			// (at prev) becomes the new idle finger.
			var cost int
			if o == free {
				cost = 0 // first placement of second finger is free
			} else {
				cost = dist(o, curr)
			}
			if v := dp[o] + cost; v < next[prev] {
				next[prev] = v
			}
		}
		dp, next = next, dp
	}

	ans := inf
	for _, v := range dp {
		if v < ans {
			ans = v
		}
	}
	return ans
}
