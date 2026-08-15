package main

// 3702. Longest Subsequence With Non-Zero Bitwise XOR
//
// Let X be the XOR of every element.
//
//   - If X != 0, the whole array already works, so the answer is len(nums).
//   - If X == 0, then for any split into kept/removed we have
//     XOR(kept) ^ XOR(removed) == 0, i.e. XOR(kept) == XOR(removed).
//     Removing a single non-zero element v therefore leaves a subsequence
//     whose XOR is v != 0, giving len(nums)-1 (and len(nums) is impossible).
//   - If X == 0 because every element is zero, no subsequence has a
//     non-zero XOR, so the answer is 0.
//
// One pass, constant extra space.
func longestSubsequence(nums []int) int {
	total := 0
	hasNonZero := false
	for _, v := range nums {
		total ^= v
		if v != 0 {
			hasNonZero = true
		}
	}

	if total != 0 {
		return len(nums)
	}
	if !hasNonZero {
		return 0
	}
	return len(nums) - 1
}
