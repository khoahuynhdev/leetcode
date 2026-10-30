package main

import "container/heap"

// Approach: Dijkstra on an expanded state graph.
// A state is (node, run) where run is the length of the current streak of
// identical labels ending at node (1..k). Moving along edge u->v either
// extends the streak (labels equal, run+1 must stay <= k) or resets it to 1.
// The first time node n-1 is popped from the heap (in any run state) we have
// the minimum total weight because all edge weights are positive.

type state struct {
	dist int
	node int
	run  int
}

type minHeap []state

func (h minHeap) Len() int            { return len(h) }
func (h minHeap) Less(i, j int) bool  { return h[i].dist < h[j].dist }
func (h minHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *minHeap) Push(x interface{}) { *h = append(*h, x.(state)) }
func (h *minHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func shortestPath(n int, edges [][]int, labels string, k int) int {
	type edge struct{ to, w int }
	adj := make([][]edge, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], edge{e[1], e[2]})
	}

	const inf = int(^uint(0) >> 1)
	// dist[node*(k+1)+run] for run in 1..k
	stride := k + 1
	dist := make([]int, n*stride)
	for i := range dist {
		dist[i] = inf
	}

	dist[0*stride+1] = 0
	h := &minHeap{{0, 0, 1}}
	for h.Len() > 0 {
		cur := heap.Pop(h).(state)
		if cur.dist > dist[cur.node*stride+cur.run] {
			continue
		}
		if cur.node == n-1 {
			return cur.dist
		}
		for _, e := range adj[cur.node] {
			run := 1
			if labels[e.to] == labels[cur.node] {
				run = cur.run + 1
				if run > k {
					continue
				}
			}
			nd := cur.dist + e.w
			idx := e.to*stride + run
			if nd < dist[idx] {
				dist[idx] = nd
				heap.Push(h, state{nd, e.to, run})
			}
		}
	}
	return -1
}
