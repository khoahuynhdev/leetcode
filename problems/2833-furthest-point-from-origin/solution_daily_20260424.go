package main

// furthestDistanceFromOrigin returns the furthest distance from the origin
// reachable after interpreting every '_' as either 'L' or 'R'.
//
// Approach: position on a number line is commutative, so order does not matter.
// Count 'L', 'R', and '_'. The fixed net displacement is |L - R|, and each '_'
// is a free move that optimally extends that displacement by 1. The answer is
// therefore |L - R| + count('_'). Runs in O(n) time and O(1) space.
func furthestDistanceFromOrigin(moves string) int {
	var left, right, free int
	for i := 0; i < len(moves); i++ {
		switch moves[i] {
		case 'L':
			left++
		case 'R':
			right++
		case '_':
			free++
		}
	}
	diff := left - right
	if diff < 0 {
		diff = -diff
	}
	return diff + free
}
