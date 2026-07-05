package main

// Problem 1301. Number of Paths with Max Score
//
// Bottom-up dynamic programming over the board. Starting from 'S' at the
// bottom-right, every legal move (up, left, up-left) goes toward 'E' at the
// top-left, so a cell (i, j) is reached from one of (i+1, j), (i, j+1) or
// (i+1, j+1). We sweep i and j from high to low so predecessors are solved
// first. For each cell we carry two values: the maximum collectable sum
// (dp) and the number of paths achieving it (cnt), with cnt taken modulo
// 1e9+7. Unreachable cells are marked with dp = -1.
func pathsWithMaxScore(board []string) []int {
	const mod = 1_000_000_007
	n := len(board)

	dp := make([][]int, n)
	cnt := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
		cnt[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = -1 // sentinel: unreachable
		}
	}

	// Base case: the path starts on 'S' at the bottom-right corner.
	dp[n-1][n-1] = 0
	cnt[n-1][n-1] = 1

	// Predecessors of (i, j): the cell we stood on before moving here.
	dirs := [3][2]int{{1, 0}, {0, 1}, {1, 1}}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if i == n-1 && j == n-1 {
				continue // start cell already initialized
			}
			if board[i][j] == 'X' {
				continue // obstacle: stays unreachable
			}

			best := -1
			ways := 0
			for _, d := range dirs {
				pi, pj := i+d[0], j+d[1]
				if pi >= n || pj >= n || dp[pi][pj] < 0 {
					continue
				}
				if dp[pi][pj] > best {
					best = dp[pi][pj]
					ways = cnt[pi][pj]
				} else if dp[pi][pj] == best {
					ways = (ways + cnt[pi][pj]) % mod
				}
			}

			if best < 0 {
				continue // no reachable predecessor
			}

			value := 0
			if board[i][j] != 'E' { // 'E' and 'S' contribute nothing
				value = int(board[i][j] - '0')
			}
			dp[i][j] = best + value
			cnt[i][j] = ways
		}
	}

	if dp[0][0] < 0 {
		return []int{0, 0}
	}
	return []int{dp[0][0], cnt[0][0]}
}
