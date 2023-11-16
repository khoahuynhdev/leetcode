package solution

import (
	"math"
	"strconv"
)

// My solution
// intuition: there are 2**n presentation and only n num in the input
// costly to generate all possible -> only need to find 1
// use a hashmap {int, bool} to store the one we've seen by looping the
// find the first value that not existed in the hashmap
// O(n) space
// O(n) time
func findDifferentBinaryString(nums []string) string {
	maxVal := math.Pow(2, float64(len(nums)))
	numDict := make(map[int]bool)
	for _, str := range nums {
		intVal, _ := strconv.ParseInt(str, 2, 64)
		numDict[int(intVal)] = true
	}
	for i := 0; i < int(maxVal); i++ {
		if !numDict[i] {
			return EnsureStrLen(strconv.FormatInt(int64(i), 2), len(nums))
		}
	}
	return ""
}

func EnsureStrLen(str string, length int) string {
	prepend := make([]byte, length-len(str))
	for i := range prepend {
		prepend[i] = '0'
	}
	return string(append(prepend, []byte(str)...))
}

// Cantor's diagonal argument is a proof in set theory.
// We start by initializing the answer ans to an empty string. To build ans,
// we need to assign either "0" or "1" to each index i for indices 0 to n - 1. How do we assign them so ans is guaranteed to be different from every string in nums? We know that two strings are different, as long as they differ by at least one character. We can intentionally construct our ans based on this fact.
// For each index i, we will check the ithi^{th}ith character of the ithi^{th}ith
// string in ans. That is, we check curr = nums[i][i]. We then assign ans[i] to the opposite of curr.
// That is, if curr = "0", we assign ans[i] = "1". If curr = "1", we assign ans[i] = "0".
// O(1) space
// O(n) time

func findDifferentBinaryStringImprove(nums []string) string {
	ans := make([]byte, len(nums))
	for i := 0; i < len(nums); i++ {
		cur := nums[i][i]
		if cur == byte('0') {
			ans[i] = byte('1')
		} else {
			ans[i] = byte('0')
		}
	}

	return string(ans)
}
