package main

func validPath(n int, edges [][]int, source int, destination int) bool {
    graph := makeGraph(edges)
    
    // Check not to visit to a node which is visited
    visited := map[int]bool{}
    
    queue := []int{source}
    for len(queue) > 0 {

        length := len(queue)
        for i := 0; i < length; i++ {
            curNode := queue[0]
            queue = queue[1:]

            // Check if there is a way from source to destination    
            if curNode == destination {
                return true
            }
            
            // Check not to add neighbours of visited node to queue
            if !visited[curNode] {
                queue = append(queue, graph[curNode]...)
            } 
            
            visited[curNode] = true
        }
    }
    return false
}

func makeGraph(edges [][]int) map[int][]int {
    graph := map[int][]int{}

    for _, edge := range edges {
        source := edge[0]
        destination := edge[1]
        graph[source] = append(graph[source], destination)
        graph[destination] = append(graph[destination], source)
    }

    return graph
}
