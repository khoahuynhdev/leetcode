package main

// closetTarget returns the shortest number of steps needed to reach `target`
// in a circular array `words` starting from `startIndex`. For every index
// where `words[i] == target`, the forward distance is `(i - startIndex + n) % n`
// and the backward distance is `n - forward`. The answer is the minimum of
// `min(forward, backward)` over all matching indices, or -1 if none match.
func closetTarget(words []string, target string, startIndex int) int {
	n := len(words)
	best := -1
	for i, w := range words {
		if w != target {
			continue
		}
		forward := (i - startIndex + n) % n
		backward := n - forward
		d := forward
		if backward < d {
			d = backward
		}
		if best == -1 || d < best {
			best = d
		}
	}
	return best
}
