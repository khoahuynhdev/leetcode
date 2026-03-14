package main

// Backtracking approach: generate happy strings in lexicographical order
// by trying 'a', 'b', 'c' at each position, skipping characters that match
// the previous one. Stop as soon as we reach the k-th string.
func getHappyString(n int, k int) string {
	var result string
	count := 0
	chars := []byte{'a', 'b', 'c'}

	var backtrack func(current []byte)
	backtrack = func(current []byte) {
		if result != "" {
			return
		}
		if len(current) == n {
			count++
			if count == k {
				result = string(current)
			}
			return
		}
		for _, c := range chars {
			if len(current) > 0 && current[len(current)-1] == c {
				continue
			}
			backtrack(append(current, c))
		}
	}

	backtrack([]byte{})
	return result
}
