package main

// minJumps returns the minimum number of steps to reach the last index of arr,
// where from index i you can move to i+1, i-1, or any j with arr[j] == arr[i].
//
// Approach: BFS on an implicit unweighted graph from index 0. Indices sharing a
// value form a "group" — we precompute these once, and on the first BFS visit to
// any member of a group we enqueue every unvisited member and then delete the
// group from the map. Deleting is the key optimization: it prevents O(n^2)
// blow-up on inputs where many indices share a single value (e.g. all 7s).
func minJumps(arr []int) int {
	n := len(arr)
	if n <= 1 {
		return 0
	}

	groups := make(map[int][]int, n)
	for i, v := range arr {
		groups[v] = append(groups[v], i)
	}

	visited := make([]bool, n)
	visited[0] = true
	frontier := []int{0}
	steps := 0

	for len(frontier) > 0 {
		next := make([]int, 0, len(frontier))
		for _, i := range frontier {
			if i == n-1 {
				return steps
			}
			if i+1 < n && !visited[i+1] {
				visited[i+1] = true
				next = append(next, i+1)
			}
			if i-1 >= 0 && !visited[i-1] {
				visited[i-1] = true
				next = append(next, i-1)
			}
			if group, ok := groups[arr[i]]; ok {
				for _, j := range group {
					if !visited[j] {
						visited[j] = true
						next = append(next, j)
					}
				}
				delete(groups, arr[i])
			}
		}
		frontier = next
		steps++
	}

	return -1
}
