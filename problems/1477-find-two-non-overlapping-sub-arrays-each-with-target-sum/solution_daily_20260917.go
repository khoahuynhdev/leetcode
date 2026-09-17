package main

// 1477. Find Two Non-overlapping Sub-arrays Each With Target Sum
//
// All elements are positive, so prefix sums are strictly increasing and a
// sliding window finds every sub-array summing to target in one pass (at most
// one per right endpoint).
//
// To pair two of them without overlap, fix the window [left..right] as the
// *second* sub-array: the first must then fit entirely in arr[0..left-1]. So we
// maintain best[i] = length of the shortest valid sub-array contained in the
// prefix arr[0..i-1], a running prefix minimum filled during the same sweep.
// The candidate answer for each window is (right-left+1) + best[left]; reading
// best[left] before folding the current window into best keeps a sub-array from
// pairing with itself.
//
// Time: O(n), Space: O(n).

// minLengthInf is the "no valid sub-array" sentinel. It is compared against
// explicitly before any addition, so it never participates in arithmetic.
const minLengthInf = 1<<31 - 1

func minSumOfLengthsDaily20260917(arr []int, target int) int {
	n := len(arr)

	// best[i] = shortest valid sub-array length within arr[0..i-1].
	best := make([]int, n+1)
	best[0] = minLengthInf

	answer := minLengthInf
	sum, left := 0, 0

	for right := 0; right < n; right++ {
		sum += arr[right]

		// Elements are positive, so shrinking strictly decreases the sum.
		for sum > target {
			sum -= arr[left]
			left++
		}

		best[right+1] = best[right]

		if sum == target {
			length := right - left + 1

			// Pair with the shortest sub-array ending before index left.
			if best[left] != minLengthInf && best[left]+length < answer {
				answer = best[left] + length
			}

			if length < best[right+1] {
				best[right+1] = length
			}
		}
	}

	if answer == minLengthInf {
		return -1
	}
	return answer
}
