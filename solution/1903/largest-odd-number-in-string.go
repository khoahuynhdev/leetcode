package solution

import (
	"strconv"
)

// https://leetcode.com/problems/largest-odd-number-in-string/description/?envType=daily-question&envId=2023-12-07

// Intuition: the largestOddNumber is the largest number from left to right ending with an odd number
// Naive approach: convert string to number -> check if it's odd
// failed at this testcase 32782489638346578713315098393010310518347382
// better approach: only check for the last char of str and convert it to number then check if it's odd
func largestOddNumber(num string) string {
	for len(num) > 0 {
		n, _ := strconv.Atoi(string(num[len(num)-1]))
		if n%2 == 1 {
			return num
		}
		num = num[:len(num)-1]
	}
	return ""
}
