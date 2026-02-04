package solution

import (
	"strings"
)

// O(n) time
// O(1) space
// intuition: minimum time means the farthest the truck must travel, the truck should
// not travel more than the required distance to collect garbage
// we can find the max index that the truck need to travel
// important points:
// - the trucks must visit each house in order
// - the trucks do not need to visit every house
// - only one garbage truck may be use at any given momment -> don't let this to
// make us loop through the trucks array
func garbageCollection(garbage []string, travel []int) int {
	ans, pIdx, mIdx, gIdx := 0, -1, -1, -1
	// find max index of each type
	for i := 0; i < len(garbage); i++ {
		if strings.Contains(garbage[i], "G") {
			gIdx = i
		}
		if strings.Contains(garbage[i], "P") {
			pIdx = i
		}
		if strings.Contains(garbage[i], "M") {
			mIdx = i
		}
	}
	pos := -1
	for i := 0; i < len(garbage); i++ {
		distTime := 0
		if pos >= 0 {
			distTime = travel[i-1]
		}
		if mIdx >= i {
			ans += distTime
			ans += strings.Count(garbage[i], "M")
		}
		if gIdx >= i {
			ans += distTime
			ans += strings.Count(garbage[i], "G")
		}
		if pIdx >= i {
			ans += distTime
			ans += strings.Count(garbage[i], "P")
		}
		pos++
	}
	return ans
}
