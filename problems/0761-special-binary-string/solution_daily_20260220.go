package main

import "sort"

// Approach: Treat 1 as '(' and 0 as ')'. A special binary string is a balanced
// parentheses sequence. Recursively find top-level special substrings, process
// their interiors, then sort them in descending order to maximize lexicographic value.
func makeLargestSpecial(s string) string {
	count := 0
	start := 0
	var subs []string

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count++
		} else {
			count--
		}
		if count == 0 {
			// Found a top-level special substring s[start..i]
			// Strip outer '1' and '0', recurse on inner, then wrap back
			inner := makeLargestSpecial(s[start+1 : i])
			subs = append(subs, "1"+inner+"0")
			start = i + 1
		}
	}

	sort.Slice(subs, func(i, j int) bool {
		return subs[i] > subs[j] // descending lexicographic order
	})

	result := ""
	for _, sub := range subs {
		result += sub
	}
	return result
}
