package main

// Process String with Special Operations II (LeetCode 3614)
//
// The final `result` can grow to 10^15 characters because of the '#'
// duplication operator, so we never build it. Instead we:
//
//  1. Forward pass: record the length of result after each operation.
//  2. Backward pass: starting from index k, undo each operation, rewriting
//     the index to the position it occupied in the previous string. When the
//     index lands on an appended letter, that letter is the answer.
//
// Runs in O(n) time and O(n) space, where n = len(s).
func processStr(s string, k int64) byte {
	n := len(s)

	// Pass 1: length of result after processing each character.
	lengths := make([]int64, n)
	var cur int64
	for i := 0; i < n; i++ {
		switch s[i] {
		case '*':
			if cur > 0 {
				cur--
			}
		case '#':
			cur *= 2
		case '%':
			// reverse: length unchanged
		default:
			cur++ // lowercase letter
		}
		lengths[i] = cur
	}

	// Out of bounds (covers the empty-result case as well).
	if k >= cur {
		return '.'
	}

	// Pass 2: trace index k back to the letter that produced it.
	pos := k
	for i := n - 1; i >= 0; i-- {
		var prev int64
		if i > 0 {
			prev = lengths[i-1]
		}
		switch s[i] {
		case '*':
			// Removed the tail; surviving indices are unchanged.
		case '#':
			// Duplicated; fold the second copy back onto the first.
			if pos >= prev {
				pos -= prev
			}
		case '%':
			// Reversed; mirror the index.
			pos = prev - 1 - pos
		default:
			// Letter appended at index prev.
			if pos == prev {
				return s[i]
			}
		}
	}

	return '.'
}
