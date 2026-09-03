package main

// 3876. Construct Uniform Parity Array II
//
// Solve the two target parities independently and take the OR.
//
// An element that already has the target parity can simply be kept, so only
// wrong-parity elements matter, and each of them must use the subtraction
// option with a partner that is both of the required parity and strictly
// smaller (the result has to be >= 1). Partners may be reused freely.
//
//   - All even: every odd element needs a strictly smaller odd partner, which
//     the smallest odd element can never have. So all-even works only when the
//     array contains no odd elements at all.
//   - All odd: every even element needs a strictly smaller odd partner, and the
//     smallest odd value is the best possible choice for all of them at once.
//     So all-odd works when there is an odd element and min(odd) < min(even).
//
// One pass tracking the two minimums is enough: O(n) time, O(1) space.
func canConstructUniformParity(nums1 []int) bool {
	const none = -1
	minOdd, minEven := none, none

	for _, v := range nums1 {
		if v%2 != 0 {
			if minOdd == none || v < minOdd {
				minOdd = v
			}
		} else {
			if minEven == none || v < minEven {
				minEven = v
			}
		}
	}

	// Only one parity present (including the empty and single-element cases):
	// keep every element as it is.
	if minOdd == none || minEven == none {
		return true
	}

	// Both parities present: all-even is out, so the smallest odd must be able
	// to repair every even element, i.e. beat the smallest even one.
	return minOdd < minEven
}
