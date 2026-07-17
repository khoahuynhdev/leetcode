package main

// 3312. Sorted GCD Pair Queries
//
// Approach: counting + number-theory sieve + prefix sum + binary search.
//
// We never build the (up to ~5e9 long) gcdPairs array. Instead we count, for
// every value g in [1, M] where M = max(nums), how many pairs have gcd exactly g.
//
//  1. cnt[v]      = frequency of value v in nums.
//  2. For each g, m = number of array elements divisible by g (sum cnt[g],
//     cnt[2g], ...). Pairs with both elements divisible by g = m*(m-1)/2. This
//     over-counts: it includes pairs whose gcd is 2g, 3g, ...
//  3. Sweep g from M down to 1 and subtract the already-finalized exact counts of
//     larger multiples: exact[g] = pairsDivisibleBy(g) - exact[2g] - exact[3g] - ...
//  4. prefix[g] = number of pairs with gcd <= g (cumulative sum of exact).
//  5. Each query index q maps to the smallest g with prefix[g] > q (binary search).
//
// Time  : O(n + M log M + Q log M)
// Space : O(M)
func gcdValues(nums []int, queries []int64) []int {
	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	cnt := make([]int, maxVal+1)
	for _, v := range nums {
		cnt[v]++
	}

	// exact[g] will hold the number of pairs whose gcd is exactly g.
	exact := make([]int64, maxVal+1)
	for g := maxVal; g >= 1; g-- {
		// m = count of array elements that are multiples of g.
		var m int64
		var over int64 // pairs already claimed by larger multiples of g
		for mult := g; mult <= maxVal; mult += g {
			m += int64(cnt[mult])
			if mult > g {
				over += exact[mult]
			}
		}
		exact[g] = m*(m-1)/2 - over
	}

	// prefix[g] = number of pairs with gcd <= g. Length maxVal+1, prefix[0] = 0.
	prefix := make([]int64, maxVal+1)
	for g := 1; g <= maxVal; g++ {
		prefix[g] = prefix[g-1] + exact[g]
	}

	answer := make([]int, len(queries))
	for i, q := range queries {
		// Smallest g in [1, maxVal] such that prefix[g] > q.
		lo, hi := 1, maxVal
		for lo < hi {
			mid := (lo + hi) / 2
			if prefix[mid] > q {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		answer[i] = lo
	}
	return answer
}
