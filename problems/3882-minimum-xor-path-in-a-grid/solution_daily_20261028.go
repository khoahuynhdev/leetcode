package main

// Approach: reachable-set DP.
// XOR is not monotonic, so keeping only the minimum prefix per cell is wrong.
// Since every value is < 1024, every prefix XOR is also < 1024, so each cell
// tracks which of the 1024 XOR values are reachable. A cell's set is the union
// of the sets above and to the left, each XORed with the cell's value.
// Rolling over columns keeps space at O(n * 1024).
// Time: O(m * n * 1024), Space: O(n * 1024).

const xorLimit = 1024

func minCost(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][xorLimit]bool, n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]
			var next [xorLimit]bool
			if i == 0 && j == 0 {
				next[v] = true
			} else {
				for x := 0; x < xorLimit; x++ {
					// dp[j] still holds the cell above; dp[j-1] is the cell to the left.
					if (i > 0 && dp[j][x]) || (j > 0 && dp[j-1][x]) {
						next[x^v] = true
					}
				}
			}
			dp[j] = next
		}
	}

	for x := 0; x < xorLimit; x++ {
		if dp[n-1][x] {
			return x
		}
	}
	return -1
}
