package main

const negInf = int64(-1e18)

func maximumTrionicSubarraySum(nums []int) int64 {
	// dp1: max sum of strictly increasing segment ending at i (at least 2 elements)
	// dp2: max sum of (increasing → decreasing) ending at i (decreasing has at least 2 elements)
	// dp3: max sum of complete trionic ending at i (second increasing has at least 2 elements)
	dp1, dp2, dp3 := negInf, negInf, negInf
	ans := negInf

	for i := 1; i < len(nums); i++ {
		curr := int64(nums[i])
		prev := int64(nums[i-1])

		// Calculate new states based on comparison with previous element
		newDp1, newDp2, newDp3 := negInf, negInf, negInf

		if curr > prev { // strictly increasing
			// dp1: extend existing increasing segment OR start fresh with [prev, curr]
			fresh := prev + curr
			if dp1 != negInf {
				newDp1 = max(dp1+curr, fresh)
			} else {
				newDp1 = fresh
			}

			// dp3: extend existing trionic OR transition from dp2 (decreasing → increasing)
			if dp3 != negInf {
				newDp3 = dp3 + curr
			}
			if dp2 != negInf {
				newDp3 = max(newDp3, dp2+curr)
			}

		} else if curr < prev { // strictly decreasing
			// dp2: extend existing decreasing OR transition from dp1 (increasing → decreasing)
			if dp2 != negInf {
				newDp2 = dp2 + curr
			}
			if dp1 != negInf {
				newDp2 = max(newDp2, dp1+curr)
			}
		}
		// if curr == prev: all states stay negInf (strict monotonicity broken)

		dp1, dp2, dp3 = newDp1, newDp2, newDp3
		ans = max(ans, dp3)
	}

	return ans
}
