package main

import "math"

// Binary search on the answer. For a candidate time T, each worker i can
// reduce height by at most x where workerTimes[i]*x*(x+1)/2 <= T.
// We sum contributions across all workers and check if they meet mountainHeight.
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	// Find minimum worker time for upper bound calculation
	minW := workerTimes[0]
	for _, w := range workerTimes {
		if w < minW {
			minW = w
		}
	}

	h := int64(mountainHeight)
	lo := int64(1)
	hi := int64(minW) * h * (h + 1) / 2

	for lo < hi {
		mid := lo + (hi-lo)/2
		if canFinish(mid, mountainHeight, workerTimes) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}

	return lo
}

func canFinish(timeLimit int64, mountainHeight int, workerTimes []int) bool {
	total := 0
	for _, w := range workerTimes {
		// Solve w * x * (x+1) / 2 <= timeLimit for x
		// x <= (-1 + sqrt(1 + 8*timeLimit/w)) / 2
		x := int((-1 + math.Sqrt(1+8*float64(timeLimit)/float64(w))) / 2)
		total += x
		if total >= mountainHeight {
			return true
		}
	}
	return false
}
