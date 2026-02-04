package solution

import (
	"fmt"
	"strconv"
)

// INTUITION:
// - use 2 pointers to check the sub string appears more than 3 times
// - Time complexity: O(n)
// - Space complexity: O(1)
func largestGoodInteger(num string) string {
	l, r := 0, 0
	max := -1
	if len(num) < 3 {
		return ""
	}
	if len(num) == 3 {
		if num[0] == num[1] && num[0] == num[2] {
			return num
		}
		return ""
	}
	for i := 0; i < len(num); i++ {
		if num[l] == num[r] {
			// bound-check r
			// fmt.Println(l, r, i)
			if r == len(num)-1 {
				if r-l >= 2 && num[l] == num[r] {
					if n, _ := strconv.Atoi(string(num[l])); max < n {
						max = n
					}
				}
				break
			}
		} else {
			if r-l >= 3 {
				if n, _ := strconv.Atoi(string(num[l])); max < n {
					max = n
				}
			}
			l = r
		}
		r++
	}
	if max > -1 {
		return fmt.Sprintf("%d%d%d", max, max, max)
	} else {
		return ""
	}
}
