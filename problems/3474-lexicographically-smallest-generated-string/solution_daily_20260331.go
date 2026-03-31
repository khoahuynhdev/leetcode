package main

// Greedy approach with constraint propagation.
// 1. Apply T constraints by stamping str2 onto the result (check overlaps).
// 2. For each F window, find the rightmost free position (last chance to break a match).
// 3. Fill free positions left-to-right with the smallest allowed character,
//    using forbidden sets from critical F windows at each position.

func generateString(str1, str2 string) string {
	n, m := len(str1), len(str2)
	L := n + m - 1
	word := make([]byte, L)
	fixed := make([]bool, L)

	// Step 1: Apply T constraints
	for i := 0; i < n; i++ {
		if str1[i] == 'T' {
			for j := 0; j < m; j++ {
				pos := i + j
				if fixed[pos] {
					if word[pos] != str2[j] {
						return ""
					}
				} else {
					word[pos] = str2[j]
					fixed[pos] = true
				}
			}
		}
	}

	// Step 2: Analyze F windows
	broken := make([]bool, n)
	rfp := make([]int, n) // rightmost free position per F window
	for i := range rfp {
		rfp[i] = -1
	}
	rfpMap := make(map[int][]int) // position -> list of F window indices with that rfp

	for i := 0; i < n; i++ {
		if str1[i] == 'F' {
			for j := 0; j < m; j++ {
				pos := i + j
				if fixed[pos] {
					if word[pos] != str2[j] {
						broken[i] = true
					}
				} else {
					rfp[i] = pos
				}
			}
			if !broken[i] {
				if rfp[i] == -1 {
					return "" // fully fixed window matches str2
				}
				rfpMap[rfp[i]] = append(rfpMap[rfp[i]], i)
			}
		}
	}

	// Step 3: Fill free positions left to right
	for j := 0; j < L; j++ {
		if fixed[j] {
			continue
		}

		// Collect forbidden chars from critical F windows ending here
		var forbiddenSet [26]bool
		if wins, ok := rfpMap[j]; ok {
			for _, i := range wins {
				if !broken[i] {
					forbiddenSet[str2[j-i]-'a'] = true
				}
			}
		}

		// Pick smallest allowed character
		var c byte
		for k := 0; k < 26; k++ {
			if !forbiddenSet[k] {
				c = byte('a' + k)
				break
			}
		}
		if c == 0 {
			return "" // all 26 chars forbidden
		}
		word[j] = c

		// Update broken status for F windows containing position j
		lo := j - m + 1
		if lo < 0 {
			lo = 0
		}
		hi := j
		if hi >= n {
			hi = n - 1
		}
		for i := lo; i <= hi; i++ {
			if str1[i] == 'F' && !broken[i] {
				if word[j] != str2[j-i] {
					broken[i] = true
				}
			}
		}
	}

	return string(word)
}
