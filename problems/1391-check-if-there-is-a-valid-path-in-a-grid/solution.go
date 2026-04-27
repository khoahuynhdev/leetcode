package main

// hasValidPath performs an iterative BFS from (0, 0) on a graph where each
// edge requires mutual street-direction agreement: cell A may step to its
// neighbor B in direction d only when A's street exposes d AND B's street
// exposes the opposite of d. Each of the six street types is encoded as the
// pair of cardinal directions it connects.
func hasValidPath(grid [][]int) bool {
	m := len(grid)
	if m == 0 {
		return false
	}
	n := len(grid[0])
	if n == 0 {
		return false
	}

	// Direction codes: 0=left, 1=right, 2=up, 3=down.
	dr := [4]int{0, 0, -1, 1}
	dc := [4]int{-1, 1, 0, 0}
	opp := [4]int{1, 0, 3, 2}

	// streets[t] = the two directions exposed by street type t (1..6).
	streets := [7][2]int{
		{}, // index 0 unused
		{0, 1}, // 1: left  - right
		{2, 3}, // 2: up    - down
		{0, 3}, // 3: left  - down
		{1, 3}, // 4: right - down
		{0, 2}, // 5: left  - up
		{1, 2}, // 6: right - up
	}

	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	visited[0][0] = true

	queue := [][2]int{{0, 0}}
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		r, c := cell[0], cell[1]

		if r == m-1 && c == n-1 {
			return true
		}

		for _, d := range streets[grid[r][c]] {
			nr, nc := r+dr[d], c+dc[d]
			if nr < 0 || nr >= m || nc < 0 || nc >= n {
				continue
			}
			if visited[nr][nc] {
				continue
			}
			ndirs := streets[grid[nr][nc]]
			if ndirs[0] != opp[d] && ndirs[1] != opp[d] {
				continue
			}
			visited[nr][nc] = true
			queue = append(queue, [2]int{nr, nc})
		}
	}

	return false
}
