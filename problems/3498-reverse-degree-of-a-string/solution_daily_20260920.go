package main

// 3498. Reverse Degree of a String
//
// Single pass over the string. For each character, its rank in the reversed
// alphabet ('a' = 26, ..., 'z' = 1) is 'z' - c + 1, and its 1-indexed position
// is i + 1. Accumulate the product of the two for every character.
//
// Time: O(n). Space: O(1).
func reverseDegree(s string) int {
	total := 0
	for i := 0; i < len(s); i++ {
		rank := int('z'-s[i]) + 1
		total += rank * (i + 1)
	}
	return total
}
