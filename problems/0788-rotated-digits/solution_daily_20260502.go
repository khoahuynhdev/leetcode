package main

// Approach: classify each digit into invalid {3,4,7}, self-rotating {0,1,8},
// or changing {2,5,6,9}. A number is "good" iff it has no invalid digit AND
// has at least one changing digit (so its rotation actually differs from x).
// We walk every x in [1, n] and check its digits — O(n log n) total.

func rotatedDigitsDaily20260502(n int) int {
	count := 0
	for x := 1; x <= n; x++ {
		if isGoodDaily20260502(x) {
			count++
		}
	}
	return count
}

func isGoodDaily20260502(x int) bool {
	hasChanging := false
	for x > 0 {
		switch x % 10 {
		case 3, 4, 7:
			return false
		case 2, 5, 6, 9:
			hasChanging = true
		}
		x /= 10
	}
	return hasChanging
}
