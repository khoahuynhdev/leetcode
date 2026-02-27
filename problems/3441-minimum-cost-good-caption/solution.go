package main

// Approach: DP with states dp[i][c][k] where i is position, c is assigned character (0-25),
// and k is the consecutive run length of c ending at i (capped at 3).
// Transitions: extend run (k->min(k+1,3)) or start new char (only if k>=3).
// Reconstruct lexicographically smallest answer by greedy forward pass.

const inf = 1<<60

func minCostGoodCaption(caption string) string {
	n := len(caption)
	if n < 3 {
		return ""
	}

	// dp[i][c][k]: min cost for caption[0..i] with position i assigned char c,
	// run length = k (k: 0->1, 1->2, 2->3+)
	// We use k in {0,1,2} representing run lengths {1,2,3+}
	dp := make([][26][3]int, n)
	parent := make([][26][3][3]int, n) // parent[i][c][k] = {prevC, prevK, -1 if no parent}

	for i := range dp {
		for c := 0; c < 26; c++ {
			for k := 0; k < 3; k++ {
				dp[i][c][k] = inf
				parent[i][c][k] = [3]int{-1, -1, -1}
			}
		}
	}

	// Base case: position 0
	for c := 0; c < 26; c++ {
		cost := abs(int(caption[0]) - ('a' + c))
		dp[0][c][0] = cost // run length 1
	}

	// Fill DP
	for i := 0; i < n-1; i++ {
		nextCost := func(c int) int {
			return abs(int(caption[i+1]) - ('a' + c))
		}
		for c := 0; c < 26; c++ {
			for k := 0; k < 3; k++ {
				if dp[i][c][k] == inf {
					continue
				}
				val := dp[i][c][k]

				// Continue same character
				nk := k + 1
				if nk > 2 {
					nk = 2
				}
				newVal := val + nextCost(c)
				if newVal < dp[i+1][c][nk] {
					dp[i+1][c][nk] = newVal
					parent[i+1][c][nk] = [3]int{c, k, i}
				}

				// Start new character (only if run >= 3, i.e., k == 2)
				if k == 2 {
					for nc := 0; nc < 26; nc++ {
						if nc == c {
							continue
						}
						newVal2 := val + nextCost(nc)
						if newVal2 < dp[i+1][nc][0] {
							dp[i+1][nc][0] = newVal2
							parent[i+1][nc][0] = [3]int{c, k, i}
						}
					}
				}
			}
		}
	}

	// Find minimum cost at position n-1 with k == 2 (run >= 3)
	bestCost := inf
	for c := 0; c < 26; c++ {
		if dp[n-1][c][2] < bestCost {
			bestCost = dp[n-1][c][2]
		}
	}
	if bestCost == inf {
		return ""
	}

	// Reconstruct: find lexicographically smallest among all optimal solutions
	// We need suffix DP for this. suffDP[i][c][k] = min cost from position i to n-1
	// given that position i is assigned char c with run length state k.
	// Then reconstruct forward greedily.

	// Compute suffix DP
	suffDP := make([][26][3]int, n)
	for i := range suffDP {
		for c := 0; c < 26; c++ {
			for k := 0; k < 3; k++ {
				suffDP[i][c][k] = inf
			}
		}
	}
	// Base: position n-1, must have k==2 for valid ending
	for c := 0; c < 26; c++ {
		suffDP[n-1][c][2] = abs(int(caption[n-1]) - ('a' + c))
	}

	// Fill suffix DP backwards
	for i := n - 2; i >= 0; i-- {
		curCost := func(c int) int {
			return abs(int(caption[i]) - ('a' + c))
		}
		for c := 0; c < 26; c++ {
			for k := 0; k < 3; k++ {
				// This state means: position i is char c, run length state is k
				// What's the min cost from i to n-1?
				cc := curCost(c)

				// Transition to i+1: continue same char
				nk := k + 1
				if nk > 2 {
					nk = 2
				}
				if suffDP[i+1][c][nk] != inf {
					val := cc + suffDP[i+1][c][nk]
					if val < suffDP[i][c][k] {
						suffDP[i][c][k] = val
					}
				}

				// Transition to i+1: start new char (only if k == 2)
				if k == 2 {
					for nc := 0; nc < 26; nc++ {
						if nc == c {
							continue
						}
						if suffDP[i+1][nc][0] != inf {
							val := cc + suffDP[i+1][nc][0]
							if val < suffDP[i][c][k] {
								suffDP[i][c][k] = val
							}
						}
					}
				}
			}
		}
	}

	// Greedy reconstruction
	result := make([]byte, n)

	// Pick first character: smallest c with suffDP[0][c][0] == bestCost
	curC, curK := -1, -1
	for c := 0; c < 26; c++ {
		if suffDP[0][c][0] == bestCost {
			curC = c
			curK = 0
			break
		}
	}
	result[0] = byte('a' + curC)

	for i := 1; i < n; i++ {
		remaining := suffDP[i-1][curC][curK] - abs(int(caption[i-1])-('a'+curC))
		found := false

		// Try all characters in order for lexicographic smallest
		for nc := 0; nc < 26; nc++ {
			if nc == curC {
				// Continue same character
				nk := curK + 1
				if nk > 2 {
					nk = 2
				}
				if suffDP[i][nc][nk] == remaining {
					result[i] = byte('a' + nc)
					curK = nk
					found = true
					break
				}
			} else {
				// Start new character (only allowed if current run >= 3)
				if curK == 2 && suffDP[i][nc][0] == remaining {
					result[i] = byte('a' + nc)
					curC = nc
					curK = 0
					found = true
					break
				}
			}
		}

		if !found {
			return ""
		}
	}

	return string(result)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
