package main

// Approach: each operation on either array reduces |nums1[i]-nums2[i]| by 1,
// so merge budgets into k = k1 + k2 and work on the diffs alone. Greedily
// lower the largest diffs first. Bucket the diffs by value and sweep from the
// maximum down, moving whole groups one level at a time; when the budget can't
// cover a full group, lower only part of it and stop.
// Time: O(n + M), Space: O(M), where M is the maximum diff.
func minSumSquareDiff(nums1 []int, nums2 []int, k1 int, k2 int) int64 {
	maxDiff := 0
	diffs := make([]int, len(nums1))
	for i := range nums1 {
		d := nums1[i] - nums2[i]
		if d < 0 {
			d = -d
		}
		diffs[i] = d
		if d > maxDiff {
			maxDiff = d
		}
	}

	cnt := make([]int64, maxDiff+1)
	for _, d := range diffs {
		cnt[d]++
	}

	k := int64(k1) + int64(k2)
	for v := maxDiff; v > 0 && k > 0; v-- {
		if k >= cnt[v] {
			k -= cnt[v]
			cnt[v-1] += cnt[v]
			cnt[v] = 0
		} else {
			cnt[v-1] += k
			cnt[v] -= k
			k = 0
		}
	}

	var sum int64
	for v := 1; v <= maxDiff; v++ {
		sum += cnt[v] * int64(v) * int64(v)
	}
	return sum
}
