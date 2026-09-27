package main

// Approach: "wormhole" traversal.
// First pair every '(' with its matching ')' using a stack of indices.
// Then walk the string: on a bracket, jump to its partner and flip the
// walking direction; on a letter, append it. Each segment inside a pair is
// thus read in reverse, and nested reversals cancel naturally. O(n) time.
func reverseParentheses(s string) string {
	n := len(s)
	pair := make([]int, n)
	stack := make([]int, 0, n)
	for i := 0; i < n; i++ {
		switch s[i] {
		case '(':
			stack = append(stack, i)
		case ')':
			j := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			pair[i], pair[j] = j, i
		}
	}

	res := make([]byte, 0, n)
	for i, d := 0, 1; i >= 0 && i < n; i += d {
		if s[i] == '(' || s[i] == ')' {
			i = pair[i]
			d = -d
		} else {
			res = append(res, s[i])
		}
	}
	return string(res)
}
