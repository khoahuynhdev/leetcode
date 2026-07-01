package main

import "container/heap"

// maximumSafenessFactor returns the maximum safeness factor of any path from
// (0,0) to (n-1,n-1), where the safeness factor of a path is the minimum
// Manhattan distance from any cell on the path to the nearest thief.
//
// Approach (two phases):
//  1. Multi-source BFS from every thief cell computes, for each cell, its
//     Manhattan distance to the nearest thief. In a 4-connected grid the BFS
//     step count equals that Manhattan distance.
//  2. A Dijkstra-style walk with a max-heap finds the "widest" path: we always
//     expand the reachable cell with the largest bottleneck value, where a
//     path's value is the minimum dist over its cells. The first time we pop
//     the destination, its recorded value is the answer.
func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)
	if n == 0 {
		return 0
	}

	// Phase 1: distance to nearest thief via multi-source BFS.
	dist := make([][]int, n)
	for i := range dist {
		dist[i] = make([]int, n)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	type cell struct{ r, c int }
	queue := make([]cell, 0, n*n)
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			if grid[r][c] == 1 {
				dist[r][c] = 0
				queue = append(queue, cell{r, c})
			}
		}
	}

	dirs := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, d := range dirs {
			nr, nc := cur.r+d[0], cur.c+d[1]
			if nr >= 0 && nr < n && nc >= 0 && nc < n && dist[nr][nc] == -1 {
				dist[nr][nc] = dist[cur.r][cur.c] + 1
				queue = append(queue, cell{nr, nc})
			}
		}
	}

	// Phase 2: widest path via a max-heap keyed by path safeness.
	safe := make([][]int, n)
	for i := range safe {
		safe[i] = make([]int, n)
		for j := range safe[i] {
			safe[i][j] = -1
		}
	}

	pq := &maxHeap{{dist[0][0], 0, 0}}
	heap.Init(pq)
	safe[0][0] = dist[0][0]

	for pq.Len() > 0 {
		item := heap.Pop(pq).(hnode)
		if item.r == n-1 && item.c == n-1 {
			return item.val
		}
		// Skip stale heap entries superseded by a better path.
		if item.val < safe[item.r][item.c] {
			continue
		}
		for _, d := range dirs {
			nr, nc := item.r+d[0], item.c+d[1]
			if nr < 0 || nr >= n || nc < 0 || nc >= n {
				continue
			}
			nv := item.val
			if dist[nr][nc] < nv {
				nv = dist[nr][nc]
			}
			if nv > safe[nr][nc] {
				safe[nr][nc] = nv
				heap.Push(pq, hnode{nv, nr, nc})
			}
		}
	}

	return 0
}

// hnode is a heap entry: val is the path safeness to reach cell (r, c).
type hnode struct{ val, r, c int }

// maxHeap orders entries so the largest safeness is popped first.
type maxHeap []hnode

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i].val > h[j].val }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x any) { *h = append(*h, x.(hnode)) }

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}
