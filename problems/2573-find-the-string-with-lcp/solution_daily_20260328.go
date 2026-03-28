package main

// Greedy construction + DP verification.
// 1. Assign characters greedily: for each unassigned position, use the next
//    letter. Propagate: if lcp[i][j] > 0, positions i and j share a character.
// 2. Verify by recomputing the LCP matrix via DP and comparing with input.
func findTheString(lcp [][]int) string {
	n := len(lcp)
	word := make([]byte, n)

	var c byte = 'a'
	for i := 0; i < n; i++ {
		if word[i] != 0 {
			continue
		}
		if c > 'z' {
			return ""
		}
		word[i] = c
		for j := i + 1; j < n; j++ {
			if lcp[i][j] > 0 {
				if word[j] != 0 && word[j] != c {
					return ""
				}
				word[j] = c
			}
		}
		c++
	}

	// Verify: recompute LCP matrix from the constructed string and compare.
	actual := make([][]int, n)
	for i := range actual {
		actual[i] = make([]int, n)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if word[i] == word[j] {
				if i+1 < n && j+1 < n {
					actual[i][j] = actual[i+1][j+1] + 1
				} else {
					actual[i][j] = 1
				}
			}
			if actual[i][j] != lcp[i][j] {
				return ""
			}
		}
	}

	return string(word)
}
