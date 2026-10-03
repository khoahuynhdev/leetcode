package main

// Approach: stack of indices with a boundary sentinel.
// The bottom of the stack always holds the index just before the current
// candidate valid substring (initially -1). Push indices of '('. On ')', pop;
// if the stack becomes empty the ')' is unmatched and becomes the new boundary,
// otherwise the valid run ending at i has length i - top.
// Time: O(n), Space: O(n).
func longestValidParentheses(s string) int {
	stack := []int{-1}
	best := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
			continue
		}
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			stack = append(stack, i)
		} else if l := i - stack[len(stack)-1]; l > best {
			best = l
		}
	}
	return best
}
