package main

// Approach: Sweep horizontal and vertical cuts using prefix sums and frequency maps.
// For each cut, check if sums are equal or can be equalized by discounting one cell.
// Connectivity constraint: in a section with h>=2 and w>=2 any cell is removable;
// in a single-row or single-column section only the two endpoints are removable.
// Time: O(m*n), Space: O(m*n)

func canPartitionGrid(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])

	totalSum := 0
	for i := range grid {
		for j := range grid[i] {
			totalSum += grid[i][j]
		}
	}

	// --- Horizontal cuts ---
	topFreq := map[int]int{}
	botFreq := map[int]int{}
	for i := range grid {
		for j := range grid[i] {
			botFreq[grid[i][j]]++
		}
	}

	topSum := 0
	for r := 0; r < m-1; r++ {
		for j := 0; j < n; j++ {
			v := grid[r][j]
			topFreq[v]++
			botFreq[v]--
			if botFreq[v] == 0 {
				delete(botFreq, v)
			}
			topSum += v
		}
		botSum := totalSum - topSum

		if topSum == botSum {
			return true
		}

		// Determine which section is larger and compute diff
		var diff int
		var freq map[int]int
		var height, rowStart, rowEnd int

		if topSum > botSum {
			diff = topSum - botSum
			freq = topFreq
			height = r + 1
			rowStart, rowEnd = 0, r
		} else {
			diff = botSum - topSum
			freq = botFreq
			height = m - r - 1
			rowStart, rowEnd = r+1, m-1
		}

		if height >= 2 && n >= 2 {
			if freq[diff] > 0 {
				return true
			}
		} else if height == 1 && n >= 2 {
			// Single row: only endpoints removable
			if grid[rowStart][0] == diff || grid[rowStart][n-1] == diff {
				return true
			}
		} else if n == 1 && height >= 2 {
			// Single column: only top and bottom removable
			if grid[rowStart][0] == diff || grid[rowEnd][0] == diff {
				return true
			}
		}
		// height==1 && n==1: single cell, cannot remove
	}

	// --- Vertical cuts ---
	leftFreq := map[int]int{}
	rightFreq := map[int]int{}
	for i := range grid {
		for j := range grid[i] {
			rightFreq[grid[i][j]]++
		}
	}

	leftSum := 0
	for c := 0; c < n-1; c++ {
		for i := 0; i < m; i++ {
			v := grid[i][c]
			leftFreq[v]++
			rightFreq[v]--
			if rightFreq[v] == 0 {
				delete(rightFreq, v)
			}
			leftSum += v
		}
		rightSum := totalSum - leftSum

		if leftSum == rightSum {
			return true
		}

		var diff int
		var freq map[int]int
		var width, colStart, colEnd int

		if leftSum > rightSum {
			diff = leftSum - rightSum
			freq = leftFreq
			width = c + 1
			colStart, colEnd = 0, c
		} else {
			diff = rightSum - leftSum
			freq = rightFreq
			width = n - c - 1
			colStart, colEnd = c+1, n-1
		}

		if m >= 2 && width >= 2 {
			if freq[diff] > 0 {
				return true
			}
		} else if m == 1 && width >= 2 {
			// Single row: only left and right endpoints removable
			if grid[0][colStart] == diff || grid[0][colEnd] == diff {
				return true
			}
		} else if width == 1 && m >= 2 {
			// Single column: only top and bottom removable
			if grid[0][colStart] == diff || grid[m-1][colStart] == diff {
				return true
			}
		}
	}

	return false
}
