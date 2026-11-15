package main

// Approach: prefix-sum parity counting.
// A subarray sum prefix[j] - prefix[i] is odd iff the two prefixes have
// different parity. Scan once, tracking how many prefixes seen so far are
// even (starting with the empty prefix) and odd; for each new prefix, add the
// count of opposite-parity prefixes to the answer.
// Time: O(n), Space: O(1).
func numOfSubarrays(arr []int) int {
	const mod = 1_000_000_007
	evenCount, oddCount := 1, 0
	prefix, ans := 0, 0
	for _, x := range arr {
		prefix = (prefix + x) % 2
		if prefix == 1 {
			ans += evenCount
			oddCount++
		} else {
			ans += oddCount
			evenCount++
		}
		ans %= mod
	}
	return ans
}
