package main

// earliestFinishTime returns the earliest moment a tourist can finish exactly one
// land ride and exactly one water ride, taken in either order.
//
// Approach: For the land-then-water plan, finish time for a pair (i, j) is
// max(landStart[i]+landDuration[i], waterStart[j]) + waterDuration[j]. For any
// fixed j, this is minimized by the i with the smallest land finish time. The
// water-then-land plan is symmetric. So we find min land finish and min water
// finish, then evaluate each plan in a single sweep. Final answer is the min of
// the two plans. Runs in O(n + m).
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

	// Land-then-water: pick the best water ride given the best land finish.
	bestLandThenWater := -1
	for j := 0; j < len(waterStartTime); j++ {
		start := minLandFinish
		if waterStartTime[j] > start {
			start = waterStartTime[j]
		}
		finish := start + waterDuration[j]
		if bestLandThenWater == -1 || finish < bestLandThenWater {
			bestLandThenWater = finish
		}
	}

	// Water-then-land: pick the best land ride given the best water finish.
	bestWaterThenLand := -1
	for i := 0; i < len(landStartTime); i++ {
		start := minWaterFinish
		if landStartTime[i] > start {
			start = landStartTime[i]
		}
		finish := start + landDuration[i]
		if bestWaterThenLand == -1 || finish < bestWaterThenLand {
			bestWaterThenLand = finish
		}
	}

	if bestLandThenWater < bestWaterThenLand {
		return bestLandThenWater
	}
	return bestWaterThenLand
}
