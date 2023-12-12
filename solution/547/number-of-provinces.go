package solution

func findCircleNum(isConnected [][]int) int {
	seen := make([]int, len(isConnected))
	provines := 0
	for i := 0; i < len(isConnected); i++ {
		if seen[i] == 1 {
			continue
		} else {
			provines++
			q := []int{i}
			for len(q) > 0 {
				size := len(q)
				for j := 0; j < size; j++ {
					if seen[q[j]] == 0 {
						seen[q[j]] = 1
						for idx, p := range isConnected[q[j]] {
							if seen[idx] == 0 && p == 1 {
								q = append(q, idx)
							}
						}
					}
				}
				q = q[size:]
			}
		}
	}

	return provines
}
