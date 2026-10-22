package main

// Approach: reverse and rotate-left generate the dihedral group, so the only
// reachable arrays are rotations of nums and rotations of reversed nums.
// Model a state as (flip, s):
//   flip = 0: a[i] = nums[(s+i) mod n]
//   flip = 1: a[i] = nums[(s-i) mod n]
// Rotate-left moves s by +1 (flip=0) or -1 (flip=1); reverse toggles flip and
// moves s by -1 (0 -> 1) or +1 (1 -> 0). Sorted means a[i] = i, so s must be
// the index of 0 in nums. Count the cheapest path to each valid target state:
//   increasing cyclic (flip=0): min(s, 2 + (n-s) mod n)
//   decreasing cyclic (flip=1): min((s+1) mod n + 1, 1 + (n-1-s) mod n)
// Time O(n), space O(1).
func minOperations(nums []int) int {
	n := len(nums)
	p := 0
	for i, v := range nums {
		if v == 0 {
			p = i
			break
		}
	}

	increasing, decreasing := true, true
	for i := 0; i < n; i++ {
		if nums[(p+i)%n] != i {
			increasing = false
		}
		if nums[((p-i)%n+n)%n] != i {
			decreasing = false
		}
	}

	best := -1
	update := func(c int) {
		if best == -1 || c < best {
			best = c
		}
	}
	if increasing {
		update(p)
		update(2 + (n-p)%n)
	}
	if decreasing {
		update((p+1)%n + 1)
		update(1 + (n-1-p)%n)
	}
	return best
}
