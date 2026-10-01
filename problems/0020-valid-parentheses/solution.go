package main

// Approach: Stack.
// Push every opening bracket. For a closing bracket, the top of the stack
// must be its matching opener; otherwise the string is invalid. The string
// is valid only if the stack is empty at the end.
// Time: O(n), Space: O(n).
func isValid(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	pairs := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		c := s[i]
		open, isClose := pairs[c]
		if !isClose {
			stack = append(stack, c)
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] != open {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}
