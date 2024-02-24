package solution
import (
  "container/heap"
)

type MinHeap []Meet

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].time < h[j].time
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Meet))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Meet struct {
	person int
	time   int
}

func findAllPeople(n int, meetings [][]int, firstPerson int) []int {
	graph := make([][]Meet, n)
	for _, meeting := range meetings {
		graph[meeting[0]] = append(graph[meeting[0]], Meet{meeting[1], meeting[2]})
		graph[meeting[1]] = append(graph[meeting[1]], Meet{meeting[0], meeting[2]})
	}
	visited := make([]bool, n)
	queue := MinHeap{{0, 0}, {firstPerson, 0}}

	for len(queue) > 0 {
		f := heap.Pop(&queue)
        if visited[f.(Meet).person] {
            continue;
        }
		visited[f.(Meet).person] = true
		for _, v := range graph[f.(Meet).person] {
			if v.time >= f.(Meet).time && !visited[v.person] {
				heap.Push(&queue, v)
			}
		}
	}

	ans := []int{}
	for i, v := range visited {
		if v {
			ans = append(ans, i)
		}
	}
	return ans
}
