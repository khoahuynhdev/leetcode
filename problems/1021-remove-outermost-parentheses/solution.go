package main

import "strings"

// Approach: track the nesting depth with a counter instead of a stack.
// A '(' that raises the depth from 0 opens a primitive, and a ')' that
// lowers the depth to 0 closes one. Skip those two characters and keep
// every other character. O(n) time, O(1) extra space besides the output.
func removeOuterParentheses(s string) string {
	var sb strings.Builder
	sb.Grow(len(s))
	depth := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			if depth > 0 {
				sb.WriteByte('(')
			}
			depth++
		} else {
			depth--
			if depth > 0 {
				sb.WriteByte(')')
			}
		}
	}
	return sb.String()
}
