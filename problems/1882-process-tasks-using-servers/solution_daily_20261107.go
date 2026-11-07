package main

import "container/heap"

// Approach: two min-heaps simulation.
//   - freeHeap orders idle servers by (weight, index).
//   - busyHeap orders working servers by (freeTime, weight, index).
//
// Task j arrives at second j. The current time is max(j, earliest free time when
// no server is idle). Before assigning, move every busy server whose freeTime <= time
// back into freeHeap. Then pop the best free server, record it, and push it into
// busyHeap with freeTime = time + tasks[j]. If no server is free, jump time forward
// to the top of busyHeap. Each server moves between heaps O(1) times per task,
// giving O((n + m) log n) overall.

type server struct {
	weight, index, freeAt int
}

type freeHeap []server

func (h freeHeap) Len() int { return len(h) }
func (h freeHeap) Less(i, j int) bool {
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	return h[i].index < h[j].index
}
func (h freeHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *freeHeap) Push(x any)   { *h = append(*h, x.(server)) }
func (h *freeHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type busyHeap []server

func (h busyHeap) Len() int { return len(h) }
func (h busyHeap) Less(i, j int) bool {
	if h[i].freeAt != h[j].freeAt {
		return h[i].freeAt < h[j].freeAt
	}
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	return h[i].index < h[j].index
}
func (h busyHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *busyHeap) Push(x any)   { *h = append(*h, x.(server)) }
func (h *busyHeap) Pop() any {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func assignTasks(servers []int, tasks []int) []int {
	free := make(freeHeap, len(servers))
	for i, w := range servers {
		free[i] = server{weight: w, index: i}
	}
	heap.Init(&free)
	busy := &busyHeap{}

	ans := make([]int, len(tasks))
	time := 0
	for j, d := range tasks {
		if time < j {
			time = j
		}
		// If nobody is idle, fast-forward to the next release time.
		if free.Len() == 0 && (*busy)[0].freeAt > time {
			time = (*busy)[0].freeAt
		}
		// Release every server that is done by now.
		for busy.Len() > 0 && (*busy)[0].freeAt <= time {
			heap.Push(&free, heap.Pop(busy).(server))
		}
		s := heap.Pop(&free).(server)
		ans[j] = s.index
		s.freeAt = time + d
		heap.Push(busy, s)
	}
	return ans
}
