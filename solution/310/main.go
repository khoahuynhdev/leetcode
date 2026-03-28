package main

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	adjList := make([][]int, n)
	for _, edge := range edges {
		adjList[edge[0]] = append(adjList[edge[0]], edge[1])
		adjList[edge[1]] = append(adjList[edge[1]], edge[0])
	}

	leaves := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(adjList[i]) == 1 {
			leaves = append(leaves, i)
		}
	}

	for n > 2 {
		newLeaves := make([]int, 0)
		for _, leaf := range leaves {
			neighbor := adjList[leaf][0]
			for i, node := range adjList[neighbor] {
				if node == leaf {
					adjList[neighbor] = append(adjList[neighbor][:i], adjList[neighbor][i+1:]...)
					break
				}
			}
			if len(adjList[neighbor]) == 1 {
				newLeaves = append(newLeaves, neighbor)
			}
			n--
		}
		leaves = newLeaves
	}

	return leaves   
}
