package main

// Approach: Compute total sum, then check every possible horizontal cut
// (after each row) and vertical cut (after each column) using a running
// prefix sum. A valid cut exists when the prefix sum equals totalSum/2.
// Time: O(m*n), Space: O(1).

func canPartitionGrid(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])

	totalSum := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			totalSum += grid[i][j]
		}
	}

	if totalSum%2 != 0 {
		return false
	}

	half := totalSum / 2

	// Check horizontal cuts (cut after row i, for i in 0..m-2)
	runningSum := 0
	for i := 0; i < m-1; i++ {
		for j := 0; j < n; j++ {
			runningSum += grid[i][j]
		}
		if runningSum == half {
			return true
		}
	}

	// Check vertical cuts (cut after column j, for j in 0..n-2)
	runningSum = 0
	for j := 0; j < n-1; j++ {
		for i := 0; i < m; i++ {
			runningSum += grid[i][j]
		}
		if runningSum == half {
			return true
		}
	}

	return false
}
