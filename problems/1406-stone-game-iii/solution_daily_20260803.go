package main

// 1406. Stone Game III
//
// Both players always take from the front of the row, so the whole remaining
// game is described by a single index i. Instead of tracking each player's
// score separately, define
//
//	dp[i] = (score of the player to move) - (score of the other player),
//	        over stoneValue[i:], with both playing optimally.
//
// Taking k in {1,2,3} stones banks their sum and hands the turn over; from
// i+k the opponent becomes the mover, so their advantage dp[i+k] counts
// against us:
//
//	dp[i] = max over k in {1,2,3}, i+k <= n of ( sum(stoneValue[i:i+k]) - dp[i+k] )
//	dp[n] = 0
//
// Each state looks at most three cells to its right, so three rolling
// variables replace the table: O(n) time, O(1) space. Alice moves first, so
// dp[0] is her final margin.
func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)

	// next1, next2, next3 hold dp[i+1], dp[i+2], dp[i+3].
	// Positions past n are never read thanks to the i+k <= n guard below.
	next1, next2, next3 := 0, 0, 0

	for i := n - 1; i >= 0; i-- {
		// k = 1 is always legal for i < n, so seed the max with it
		// rather than 0 (dp values may legitimately be negative).
		sum := stoneValue[i]
		best := sum - next1

		if i+2 <= n {
			sum += stoneValue[i+1]
			if cand := sum - next2; cand > best {
				best = cand
			}
		}
		if i+3 <= n {
			sum += stoneValue[i+2]
			if cand := sum - next3; cand > best {
				best = cand
			}
		}

		next1, next2, next3 = best, next1, next2
	}

	// After the loop next1 == dp[0] (or 0 for an empty row).
	switch {
	case next1 > 0:
		return "Alice"
	case next1 < 0:
		return "Bob"
	default:
		return "Tie"
	}
}
