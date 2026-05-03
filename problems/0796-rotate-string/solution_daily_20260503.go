package main

import "strings"

// rotateString reports whether s can become goal after some number of left
// shifts. Every rotation of s appears as a length-len(s) substring of s+s,
// so the problem reduces to a length check plus one substring search.
func rotateString(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	return strings.Contains(s+s, goal)
}
