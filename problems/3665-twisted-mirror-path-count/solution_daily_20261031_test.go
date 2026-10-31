package main

import (
	"math/rand"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name     string
		grid     [][]int
		expected int
	}{
		{
			"example 1: mixed mirrors give 5 paths",
			[][]int{{0, 1, 0}, {0, 0, 1}, {1, 0, 0}},
			5,
		},
		{
			"example 2: 2x2 no mirrors",
			[][]int{{0, 0}, {0, 0}},
			2,
		},
		{
			"example 3: chained reflections, one path goes out of bounds",
			[][]int{{0, 1, 1}, {1, 1, 0}},
			1,
		},
		{
			"edge case: 3x3 no mirrors equals binomial C(4,2)",
			[][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			6,
		},
		{
			"edge case: both neighbours of start are mirrors deflecting onto target",
			[][]int{{0, 1}, {1, 0}},
			2,
		},
		{
			"edge case: mirrors adjacent to target deflect into it",
			[][]int{{0, 0, 0}, {0, 0, 1}, {0, 1, 0}},
			4,
		},
		{
			"edge case: top row of mirrors forces robot into bottom row",
			[][]int{{0, 1, 1, 1}, {0, 0, 0, 0}},
			2,
		},
		{
			"edge case: double reflection survives while the other branch exits the grid",
			[][]int{{0, 1}, {1, 1}, {1, 0}},
			1,
		},
		{
			"edge case: large empty grid requires modulo",
			emptyGrid(500, 500),
			binomMod(998, 499),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := uniquePaths(tt.grid)
			if got != tt.expected {
				t.Errorf("got %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestUniquePathsBruteForce cross-checks the DP against a direct simulation
// on small random grids.
func TestUniquePathsBruteForce(t *testing.T) {
	r := rand.New(rand.NewSource(3665))
	for iter := 0; iter < 500; iter++ {
		m, n := 2+r.Intn(4), 2+r.Intn(4)
		grid := make([][]int, m)
		for i := range grid {
			grid[i] = make([]int, n)
			for j := range grid[i] {
				if r.Intn(3) == 0 {
					grid[i][j] = 1
				}
			}
		}
		grid[0][0], grid[m-1][n-1] = 0, 0
		want := bruteForce(grid)
		if got := uniquePaths(grid); got != want {
			t.Fatalf("grid %v: got %d, want %d", grid, got, want)
		}
	}
}

func bruteForce(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// move simulates stepping in direction d (0 = right, 1 = down) from (i, j),
	// following reflections until landing on an empty cell, the target, or out of bounds.
	var dfs func(i, j int) int
	move := func(i, j, d int) (int, int, bool) {
		for {
			if d == 0 {
				j++
			} else {
				i++
			}
			if i >= m || j >= n {
				return 0, 0, false
			}
			if grid[i][j] == 0 {
				return i, j, true
			}
			d ^= 1
		}
	}
	dfs = func(i, j int) int {
		if i == m-1 && j == n-1 {
			return 1
		}
		total := 0
		for d := 0; d < 2; d++ {
			if ni, nj, ok := move(i, j, d); ok {
				total += dfs(ni, nj)
			}
		}
		return total
	}
	return dfs(0, 0)
}

func emptyGrid(m, n int) [][]int {
	g := make([][]int, m)
	for i := range g {
		g[i] = make([]int, n)
	}
	return g
}

func binomMod(n, k int) int {
	row := make([]int, k+1)
	row[0] = 1
	for i := 1; i <= n; i++ {
		for j := min(i, k); j > 0; j-- {
			row[j] = (row[j] + row[j-1]) % mod
		}
	}
	return row[k]
}
