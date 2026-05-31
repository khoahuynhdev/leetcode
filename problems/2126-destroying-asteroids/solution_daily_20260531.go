package main

// Greedy: sort asteroids ascending, then absorb them in order.
// If the planet ever meets an asteroid larger than its current mass,
// no ordering can save it because every remaining asteroid is at least
// that size. Using int (64-bit) avoids overflow as cumulative mass
// can reach ~10^10.

import "sort"

func asteroidsDestroyed(mass int, asteroids []int) bool {
	sort.Ints(asteroids)
	planet := mass
	for _, a := range asteroids {
		if planet < a {
			return false
		}
		planet += a
	}
	return true
}
