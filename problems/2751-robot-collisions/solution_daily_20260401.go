// Robot Collisions: Sort robots by position, then use a stack to simulate
// collisions between right-moving and left-moving robots.
package main

import "sort"

func survivedRobotsHealths(positions []int, healths []int, directions string) []int {
	n := len(positions)

	// Create indices sorted by position
	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}
	sort.Slice(indices, func(a, b int) bool {
		return positions[indices[a]] < positions[indices[b]]
	})

	// Stack holds indices of right-moving robots
	stack := []int{}

	// Track which robots are removed
	removed := make([]bool, n)

	for _, i := range indices {
		if directions[i] == 'R' {
			stack = append(stack, i)
			continue
		}

		// directions[i] == 'L': resolve collisions with right-movers on stack
		for len(stack) > 0 && !removed[i] {
			top := stack[len(stack)-1]
			if healths[top] < healths[i] {
				// Right-mover destroyed
				removed[top] = true
				stack = stack[:len(stack)-1]
				healths[i]--
			} else if healths[top] == healths[i] {
				// Both destroyed
				removed[top] = true
				stack = stack[:len(stack)-1]
				removed[i] = true
			} else {
				// Left-mover destroyed
				healths[top]--
				removed[i] = true
			}
		}
	}

	// Collect survivors in original order
	result := []int{}
	for i := 0; i < n; i++ {
		if !removed[i] {
			result = append(result, healths[i])
		}
	}
	return result
}
