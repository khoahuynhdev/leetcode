package main

// Approach: grid DP with an extra "incoming direction" dimension.
//
// dp[i][j][0] = number of ways to arrive at (i, j) while moving RIGHT
// dp[i][j][1] = number of ways to arrive at (i, j) while moving DOWN
//
// Every move goes right or down, so processing cells in row-major order
// guarantees all contributions to a cell are final before it is expanded.
//
//   - Empty cell: the robot may choose either direction, so all arrivals
//     (dp0 + dp1) flow right into (i, j+1) as RIGHT and down into (i+1, j) as DOWN.
//   - Mirror cell: the robot never "stops" here. Arriving RIGHT deflects it DOWN
//     into (i+1, j); arriving DOWN deflects it RIGHT into (i, j+1). Chains of
//     mirrors are handled naturally because the deflected arrival is just
//     another arrival that will be processed later.
//
// Moves that leave the grid are simply dropped (invalid paths).
// Time: O(m*n), Space: O(n) using two rolling rows.

const mod = 1_000_000_007

func uniquePaths(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// cur[j] holds arrivals for the current row, next[j] for the row below.
	cur := make([][2]int, n)
	next := make([][2]int, n)
	cur[0][0] = 1 // treat the start as an arrival (direction irrelevant: cell is empty)

	for i := 0; i < m; i++ {
		for j := range next {
			next[j] = [2]int{}
		}
		for j := 0; j < n; j++ {
			right, down := cur[j][0], cur[j][1]
			if right == 0 && down == 0 {
				continue
			}
			if grid[i][j] == 0 {
				total := (right + down) % mod
				if i == m-1 && j == n-1 {
					return total
				}
				if j+1 < n {
					cur[j+1][0] = (cur[j+1][0] + total) % mod
				}
				if i+1 < m {
					next[j][1] = (next[j][1] + total) % mod
				}
			} else {
				// Arrived moving right -> deflected down.
				if right > 0 && i+1 < m {
					next[j][1] = (next[j][1] + right) % mod
				}
				// Arrived moving down -> deflected right.
				if down > 0 && j+1 < n {
					cur[j+1][0] = (cur[j+1][0] + down) % mod
				}
			}
		}
		cur, next = next, cur
	}
	return 0
}
