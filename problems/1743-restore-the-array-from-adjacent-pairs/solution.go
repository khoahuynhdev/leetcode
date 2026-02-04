package solution

import (
	"fmt"
)

func restoreArray(adjacentPairs [][]int) []int {
	// build the graph
	graph := make(map[int][]int)

	for _, nums := range adjacentPairs {
		graph[nums[0]] = append(graph[nums[0]], nums[1])
		graph[nums[1]] = append(graph[nums[1]], nums[0])
	}

	tail := adjacentPairs[0][0] // always available
	// detect the tail
	// not so good the option
	for k := range graph {
		if len(graph[k]) == 1 {
			tail = k
			break
		}
	}

	// fmt.Println(graph)
	// fmt.Println(tail)
	// NOTE: use seen for less time when working on web editor
	visited := make(map[int]bool)
	res := []int{}
	queue := []int{tail}
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			cur := queue[i]
			if !visited[cur] {
				res = append(res, cur)
				visited[cur] = true
				for _, num := range graph[cur] {
					if !visited[num] {
						queue = append(queue, num)
					}
				}
			}
		}
		// delete value from queue
		queue = queue[size:]
	}
	return res
}

// there is a solution that uses DFS
// basically check for visited node and add them to the answer

func main() {
	// fmt.Println(restoreArray([][]int{[]int{2, 1}, []int{3, 4}, []int{3, 2}}))
	fmt.Println(restoreArray([][]int{{4, -10}, {-1, 3}, {4, -3}, {-3, 3}}))
}
