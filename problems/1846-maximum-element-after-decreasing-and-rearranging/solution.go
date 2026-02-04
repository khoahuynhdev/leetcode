package solution

import (
	"sort"
)

func maximumElementAfterDecrementingAndRearranging(arr []int) int {
	if len(arr) == 1 {
		return 1
	}
	sort.Ints(arr)
	wd := []int{arr[0], arr[1]}
	if wd[0] != 1 {
		wd[0] = 1
	}
	if wd[1]-wd[0] > 1 {
		wd[1] = wd[0] + 1
	}
	for i := len(wd); i < len(arr); i++ {
		cur := arr[i]
		if cur == wd[1] {
			wd[0] = wd[1]
			wd[1] = cur
			continue
		}
		if cur-wd[1] > 1 {
			wd[0] = wd[1]
			wd[1] = wd[1] + 1
		} else {
			wd[0] = wd[1]
			wd[1] = cur
		}

	}
	return wd[1]
}
