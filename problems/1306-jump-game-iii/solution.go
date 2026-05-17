package main

// canReach treats the array as an implicit graph and runs DFS from start.
// At each index i, the two outgoing edges are i+arr[i] and i-arr[i]. A
// visited array bounds the search to O(n) and prevents cycles between
// indices whose jumps point at each other.
func canReach(arr []int, start int) bool {
	n := len(arr)
	if n == 0 {
		return false
	}
	visited := make([]bool, n)

	var dfs func(i int) bool
	dfs = func(i int) bool {
		if i < 0 || i >= n || visited[i] {
			return false
		}
		if arr[i] == 0 {
			return true
		}
		visited[i] = true
		return dfs(i+arr[i]) || dfs(i-arr[i])
	}

	return dfs(start)
}
