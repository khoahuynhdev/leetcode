package main

// Earliest Finish Time for Land and Water Rides II
//
// A plan runs one land ride and one water ride back-to-back in one of two
// orders. The second ride's finish time depends on the first ride only through
// the first ride's finish time, and is monotonically non-decreasing in it.
// So for each ordering we only need the earliest-finishing ride of the first
// category, then sweep the second category once. This collapses the naive
// O(n*m) pairing into an O(n+m) greedy.
func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	minLandFinish := landStartTime[0] + landDuration[0]
	for i := 1; i < len(landStartTime); i++ {
		if f := landStartTime[i] + landDuration[i]; f < minLandFinish {
			minLandFinish = f
		}
	}

	minWaterFinish := waterStartTime[0] + waterDuration[0]
	for j := 1; j < len(waterStartTime); j++ {
		if f := waterStartTime[j] + waterDuration[j]; f < minWaterFinish {
			minWaterFinish = f
		}
	}

	best := -1

	// Order 1: land ride first, then water ride j.
	for j := range waterStartTime {
		finish := max(minLandFinish, waterStartTime[j]) + waterDuration[j]
		if best == -1 || finish < best {
			best = finish
		}
	}

	// Order 2: water ride first, then land ride i.
	for i := range landStartTime {
		finish := max(minWaterFinish, landStartTime[i]) + landDuration[i]
		if finish < best {
			best = finish
		}
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
