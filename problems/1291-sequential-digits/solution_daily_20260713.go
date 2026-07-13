package main

import "sort"

// Sequential Digits (LeetCode 1291)
//
// Every number with sequential digits is a contiguous substring of the fixed
// string "123456789". We slide a window of each possible length (2..9) across
// that string, turn each window into an integer, and keep the ones that fall
// inside [low, high]. This enumerates the entire (tiny, constant-size) universe
// of candidates instead of scanning the potentially huge [low, high] range.
func sequentialDigits(low int, high int) []int {
	const digits = "123456789"

	result := []int{}
	for length := 2; length <= len(digits); length++ {
		for start := 0; start+length <= len(digits); start++ {
			num := 0
			for i := start; i < start+length; i++ {
				num = num*10 + int(digits[i]-'0')
			}
			if num >= low && num <= high {
				result = append(result, num)
			}
		}
	}

	// The generation order already yields sorted output, but sort defensively.
	sort.Ints(result)
	return result
}
