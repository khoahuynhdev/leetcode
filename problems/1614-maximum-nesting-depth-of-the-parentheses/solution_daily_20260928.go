package main

// Approach: scan once with a counter instead of a stack. Increment on '(',
// decrement on ')', and track the maximum counter value seen. Since s is a
// valid parentheses string, the counter always equals the current nesting depth.
// Time: O(n), Space: O(1).
func maxDepth(s string) int {
	depth, best := 0, 0
	for _, c := range s {
		switch c {
		case '(':
			depth++
			if depth > best {
				best = depth
			}
		case ')':
			depth--
		}
	}
	return best
}
