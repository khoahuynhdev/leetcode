// 3568. Minimum Moves to Clean the Classroom
//
// BFS over the state space (row, col, collectedMask, remainingEnergy).
// A plain grid BFS is not enough: the answer depends on which litter has
// already been picked up and on how much energy is left, so both are folded
// into the state. Since at most 10 'L' cells exist, the set of collected
// litter fits in a 10-bit mask.
//
// Every move costs exactly 1, so BFS explores states in order of move count
// and the first state whose mask is full is optimal. To keep the state space
// small we do not treat energy as a separate BFS dimension; instead we record,
// for each (cell, mask), the greatest remaining energy we have ever reached it
// with. Arriving again with less-or-equal energy is strictly dominated: any
// continuation from the weaker state is also available from the stronger one.

package main

type classroomState struct {
	r, c, mask, energy int
}

func minMoves(classroom []string, energy int) int {
	m := len(classroom)
	if m == 0 {
		return -1
	}
	n := len(classroom[0])

	// litterID maps a flattened cell index to its bit position, or -1.
	litterID := make([]int, m*n)
	for i := range litterID {
		litterID[i] = -1
	}
	total := 0
	sr, sc := -1, -1
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			switch classroom[r][c] {
			case 'S':
				sr, sc = r, c
			case 'L':
				litterID[r*n+c] = total
				total++
			}
		}
	}
	if sr < 0 {
		return -1
	}

	full := (1 << total) - 1
	if full == 0 {
		return 0 // nothing to clean
	}

	// best[cell][mask] is the greatest remaining energy with which that pair
	// has already been reached; -1 means it has never been reached.
	best := make([][]int, m*n)
	for i := range best {
		best[i] = make([]int, full+1)
		for j := range best[i] {
			best[i][j] = -1
		}
	}
	best[sr*n+sc][0] = energy

	dirs := [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	queue := []classroomState{{r: sr, c: sc, mask: 0, energy: energy}}

	for moves := 0; len(queue) > 0; moves++ {
		var next []classroomState
		for _, st := range queue {
			if st.energy == 0 {
				continue // out of fuel and not standing on a reset area
			}
			for _, d := range dirs {
				nr, nc := st.r+d[0], st.c+d[1]
				if nr < 0 || nr >= m || nc < 0 || nc >= n {
					continue
				}
				ch := classroom[nr][nc]
				if ch == 'X' {
					continue
				}

				nEnergy := st.energy - 1
				nMask := st.mask
				if ch == 'R' {
					nEnergy = energy // reset areas refill regardless of level
				} else if id := litterID[nr*n+nc]; id >= 0 {
					nMask |= 1 << id
				}
				if nMask == full {
					return moves + 1
				}

				idx := nr*n + nc
				if best[idx][nMask] >= nEnergy {
					continue
				}
				best[idx][nMask] = nEnergy
				next = append(next, classroomState{r: nr, c: nc, mask: nMask, energy: nEnergy})
			}
		}
		queue = next
	}

	return -1
}
