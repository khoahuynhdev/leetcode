package main

// 3691. Maximum Total Subarray Value II
//
// Approach: The answer is the sum of the k largest subarray "values", where the
// value of nums[l..r] is max - min (all values are non-negative and subarrays
// must be distinct, so distinctness never forces a smaller pick).
//
// We binary search a cutoff range value T. For a threshold X we compute, in a
// single O(n) sliding-window pass, both the number of subarrays whose range
// (max-min) is <= X and the sum of those ranges. Two monotonic deques track the
// window max and min; each deque element owns a contiguous block of left
// endpoints (its segment length), which lets us maintain the windowed sum of
// subarray-maxes and subarray-mins in O(1) per step. From count(<=X)/sum(<=X)
// we derive count(>=T)/sum(>=T) and assemble the top-k total.
//
// Time: O(n log V), Space: O(n).
func maxTotalValue(nums []int, k int) int64 {
	n := len(nums)
	total := int64(n) * int64(n+1) / 2

	type seg struct {
		idx int
		ln  int64
	}

	// rangeCountSumLE returns (cnt, sum): the number of subarrays whose
	// (max-min) <= X, and the total of their (max-min) values.
	rangeCountSumLE := func(X int64) (int64, int64) {
		var maxDq, minDq []seg
		var sumMax, sumMin int64
		var cnt, sum int64
		left := 0
		for r := 0; r < n; r++ {
			v := nums[r]

			// Extend the max deque: r absorbs every back segment it dominates.
			var pLen int64 = 0
			for len(maxDq) > 0 && nums[maxDq[len(maxDq)-1].idx] <= v {
				e := maxDq[len(maxDq)-1]
				maxDq = maxDq[:len(maxDq)-1]
				sumMax -= int64(nums[e.idx]) * e.ln
				pLen += e.ln
			}
			maxDq = append(maxDq, seg{r, pLen + 1})
			sumMax += int64(v) * (pLen + 1)

			// Extend the min deque symmetrically.
			pLen = 0
			for len(minDq) > 0 && nums[minDq[len(minDq)-1].idx] >= v {
				e := minDq[len(minDq)-1]
				minDq = minDq[:len(minDq)-1]
				sumMin -= int64(nums[e.idx]) * e.ln
				pLen += e.ln
			}
			minDq = append(minDq, seg{r, pLen + 1})
			sumMin += int64(v) * (pLen + 1)

			// Shrink from the left while the window range exceeds X.
			for int64(nums[maxDq[0].idx]-nums[minDq[0].idx]) > X {
				sumMax -= int64(nums[maxDq[0].idx])
				maxDq[0].ln--
				if maxDq[0].ln == 0 {
					maxDq = maxDq[1:]
				}
				sumMin -= int64(nums[minDq[0].idx])
				minDq[0].ln--
				if minDq[0].ln == 0 {
					minDq = minDq[1:]
				}
				left++
			}

			cnt += int64(r - left + 1)
			sum += sumMax - sumMin
		}
		return cnt, sum
	}

	// totalSum: a window that never shrinks (X larger than any possible range).
	_, totalSum := rangeCountSumLE(int64(1) << 62)

	// geCount(T): number of subarrays with range >= T.
	geCount := func(T int64) int64 {
		if T <= 0 {
			return total
		}
		c, _ := rangeCountSumLE(T - 1)
		return total - c
	}

	mx, mn := nums[0], nums[0]
	for _, x := range nums {
		if x > mx {
			mx = x
		}
		if x < mn {
			mn = x
		}
	}

	// Largest T with geCount(T) >= k.
	lo, hi := int64(0), int64(mx-mn)
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if geCount(mid) >= int64(k) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	tStar := lo

	cLE, sLE := rangeCountSumLE(tStar) // subarrays with range <= tStar
	c1 := total - cLE                  // subarrays with range >= tStar+1
	sumGE := totalSum - sLE            // sum of ranges >= tStar+1

	return sumGE + tStar*(int64(k)-c1)
}
