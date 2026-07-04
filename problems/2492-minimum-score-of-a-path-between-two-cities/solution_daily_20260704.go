package main

// Problem 2492: Minimum Score of a Path Between Two Cities.
//
// Because a path may repeat roads and revisit cities without limit, any road in
// the connected component containing city 1 can be included in a walk from 1 to n
// (which are guaranteed connected). The best achievable score is therefore the
// minimum edge distance across every road in that component.
//
// Approach: build an adjacency list, BFS the component starting at city 1, and
// track the minimum edge weight encountered. Runs in O(n + m).
func minScore(n int, roads [][]int) int {
	type edge struct {
		to, dist int
	}

	adj := make([][]edge, n+1)
	for _, r := range roads {
		a, b, d := r[0], r[1], r[2]
		adj[a] = append(adj[a], edge{b, d})
		adj[b] = append(adj[b], edge{a, d})
	}

	const maxDist = 1<<31 - 1
	answer := maxDist

	visited := make([]bool, n+1)
	queue := []int{1}
	visited[1] = true

	for len(queue) > 0 {
		city := queue[0]
		queue = queue[1:]

		for _, e := range adj[city] {
			if e.dist < answer {
				answer = e.dist
			}
			if !visited[e.to] {
				visited[e.to] = true
				queue = append(queue, e.to)
			}
		}
	}

	return answer
}
