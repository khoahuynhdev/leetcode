package solution

// Cantor's diagonal argument: build a string that differs from nums[i] at position i.
// This guarantees the result is not equal to any string in nums.
// Time: O(n), Space: O(n)
func findDifferentBinaryStringDaily(nums []string) string {
	result := make([]byte, len(nums))
	for i, s := range nums {
		if s[i] == '0' {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	return string(result)
}
