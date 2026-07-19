package main

// Smallest Subsequence of Distinct Characters (same as LeetCode 316).
//
// Approach: monotonic-stack greedy. We build the answer on a stack, keeping it
// as close to increasing as possible. For each character we pop any greater
// character from the top of the stack as long as that character is guaranteed
// to appear again later (so we can safely place it further right). Characters
// already on the stack are skipped, ensuring each distinct char appears once.
// Each character is pushed and popped at most once, giving O(n) time.
func smallestSubsequence(s string) string {
	// last[c] holds the highest index at which byte c-'a' occurs in s.
	var last [26]int
	for i := 0; i < len(s); i++ {
		last[s[i]-'a'] = i
	}

	var inStack [26]bool
	stack := make([]byte, 0, 26)

	for i := 0; i < len(s); i++ {
		c := s[i]
		if inStack[c-'a'] {
			// Already placed; keeping one copy is enough.
			continue
		}
		// Pop larger characters that still appear later in the string.
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			if top > c && last[top-'a'] > i {
				stack = stack[:len(stack)-1]
				inStack[top-'a'] = false
			} else {
				break
			}
		}
		stack = append(stack, c)
		inStack[c-'a'] = true
	}

	return string(stack)
}
