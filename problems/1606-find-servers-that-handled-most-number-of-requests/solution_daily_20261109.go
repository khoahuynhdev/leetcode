package main

// Approach: simulate with two min-heaps.
//
//   - busy: (endTime, serverID) for servers currently working, ordered by endTime.
//   - available: free servers, but stored under a "virtual index" so that a plain
//     min-heap answers "the first free server at or after i % k, wrapping around".
//
// When a server j becomes free right before request i, we push the smallest
// value v >= i with v ≡ j (mod k), i.e. v = i + ((j-i)%k + k) % k. Every value
// in the available heap is therefore >= i, and the minimum value corresponds to
// the first free server reached by scanning i%k, i%k+1, ... with wrap-around.
// The server ID is recovered as v % k.
//
// Invariant: values in the heap are always >= the current request index. At
// step i at most one value equals i (residues are distinct per server), and it
// is the minimum, so it gets popped at step i — keeping the invariant for i+1.
//
// Time: O(n log k), Space: O(k).

import "container/heap"

type busyServer struct {
	end int
	id  int
}

type busyHeap []busyServer

func (h busyHeap) Len() int            { return len(h) }
func (h busyHeap) Less(i, j int) bool  { return h[i].end < h[j].end }
func (h busyHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *busyHeap) Push(x interface{}) { *h = append(*h, x.(busyServer)) }
func (h *busyHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

type intHeap []int

func (h intHeap) Len() int            { return len(h) }
func (h intHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *intHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *intHeap) Pop() interface{} {
	old := *h
	x := old[len(old)-1]
	*h = old[:len(old)-1]
	return x
}

func busiestServers(k int, arrival []int, load []int) []int {
	count := make([]int, k)

	available := make(intHeap, k)
	for j := 0; j < k; j++ {
		available[j] = j // already a valid heap (sorted)
	}
	busy := &busyHeap{}

	for i, t := range arrival {
		// Release every server that finished at or before time t.
		for busy.Len() > 0 && (*busy)[0].end <= t {
			s := heap.Pop(busy).(busyServer)
			v := i + ((s.id-i)%k+k)%k
			heap.Push(&available, v)
		}

		if available.Len() == 0 {
			continue // all servers busy: drop the request
		}

		id := heap.Pop(&available).(int) % k
		count[id]++
		heap.Push(busy, busyServer{end: t + load[i], id: id})
	}

	maxCount := 0
	for _, c := range count {
		if c > maxCount {
			maxCount = c
		}
	}

	result := []int{}
	for id, c := range count {
		if c == maxCount {
			result = append(result, id)
		}
	}
	return result
}
