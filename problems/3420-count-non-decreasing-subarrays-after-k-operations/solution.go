package main

// Approach: sliding window scanned right-to-left with a monotonic deque.
//
// For a fixed subarray [l, r], the minimum number of operations needed to make
// it non-decreasing is sum_{i=l..r} (max(nums[l..i]) - nums[i]): every element
// must be raised to the running prefix maximum.
//
// We move the left end l from n-1 down to 0. The window is grouped into
// segments that share the same prefix maximum (a "block" led by a dominating
// element). Segment maxima strictly increase from left to right. The structure
// behaves as a monotonic stack on the left (where new elements arrive) and a
// queue on the right (where elements leave):
//
//   - Adding nums[l] on the left: every segment whose max <= nums[l] is now
//     dominated by nums[l]; each of its elements needs (nums[l] - oldMax) more
//     operations. Merge those segments into the new one.
//   - While cost > k, drop nums[r] from the right: its contribution is
//     (max of the rightmost segment) - nums[r]. Shrink that segment.
//
// For every l, all subarrays [l, l..r] are valid, adding r - l + 1 to the
// answer. Each index is pushed and popped at most once: O(n) overall.

type segment struct {
	max   int // prefix maximum shared by every element in this segment
	count int // number of window elements in this segment
}

func countNonDecreasingSubarrays(nums []int, k int) int64 {
	n := len(nums)
	// segs[head:] holds the window's segments; segs[head] is the rightmost
	// segment and segs[len-1] is the leftmost (most recently added).
	segs := make([]segment, 0, n)
	head := 0
	cost := 0
	var result int64
	r := n - 1

	for l := n - 1; l >= 0; l-- {
		cur := segment{max: nums[l], count: 1}
		for len(segs) > head && segs[len(segs)-1].max <= nums[l] {
			top := segs[len(segs)-1]
			segs = segs[:len(segs)-1]
			cost += (nums[l] - top.max) * top.count
			cur.count += top.count
		}
		segs = append(segs, cur)

		for cost > k {
			cost -= segs[head].max - nums[r]
			segs[head].count--
			if segs[head].count == 0 {
				head++
			}
			r--
		}

		result += int64(r - l + 1)
	}
	return result
}
