package solution

func numBusesToDestination(routes [][]int, source int, target int) int {
	g := make(map[int][]int)
	if source == target {
		return 0
	}
	for bus, stops := range routes {
		src := false
		tar := false
		// bus connects to all routes
		for _, stop := range stops {
			if stop == source {
				src = true
			}
			if stop == target {
				tar = true
			}
			g[stop] = append(g[stop], bus)
		}
		if src && tar {
			return 1
		}
	}
	if _, ok := g[target]; !ok {
		return -1
	}
	// fmt.Println(g)

	seen := make(map[int]bool)
	queue := []int{g[target][0]}
	buses := 0
	for len(queue) > 0 {
		size := len(queue)
		buses++
		for i := 0; i < size; i++ {
			bus := queue[i]
			if !seen[bus] {
				seen[bus] = true
				for k := 0; k < len(routes[bus]); k++ {
					route := routes[bus][k]
					if route == source {
						return buses
					}
					for j := 0; j < len(g[route]); j++ {
						if !seen[g[route][j]] {
							queue = append(queue, g[route][j])
						}
					}

				}
			}
		}
		queue = queue[size:]
	}
	return -1
}
