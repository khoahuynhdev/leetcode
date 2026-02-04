	// // TODO: Implement the solution
	// // Hint: Use dynamic programming with state tracking
	// // Track three states: first increasing, decreasing, final increasing
	// stSum, ndSum, rdSum := int64(0), int64(0), int64(0)
	// end := len(nums)
	// l, p, q := nums[0], math.MaxInt32, math.MinInt32
	// lUpdated, pUpdated, qUpdated := false, false, false
	//
	// // may need to add first l to the stSum
	// // [0,-2,-1,-3,0,2,-1]
	// for i := 1; i < end; i++ {
	// 	fmt.Println("i == At index", i, "with value", nums[i], "current l:", l, "p:", p, "q:", q)
	//
	// 	if l == nums[i] {
	// 		// fmt.Println("It's equal to nums[i], resetting l")
	// 		// this is not the trionic subarray, reset
	// 		l = nums[i]
	// 		stSum = 0
	// 		continue
	// 	}
	//
	// 	for i < end && l < nums[i] {
	// 		stSum += int64(l)
	// 		l = nums[i]
	// 		lUpdated = true
	// 		i++
	// 	}
	// 	fmt.Println("Exited first loop at index", i, "current l:", l, "stSum:", stSum)
	//
	// 	if l > nums[i] {
	// 		// which condition to reset l pos and val here?
	// 		if !lUpdated {
	// 			fmt.Println("Resetting l at index", i, "with value", nums[i])
	// 			// p has not been set yet so this indicate a reset
	// 			l = nums[i]
	// 			lUpdated = false
	// 			continue
	// 		}
	// 		// maybe it the second state
	// 		fmt.Println("Found potential second state starting at index", i, "with value", nums[i], "with l = ", l)
	// 		p = l
	// 		for i < end && p > nums[i] {
	// 			// can be second state
	// 			ndSum += int64(p)
	// 			p = nums[i]
	// 			pUpdated = true
	// 			i++
	// 		}
	// 		fmt.Println("Exited second loop at index", i, "with value", nums[i], " l:", l, " p: ", p, "ndSum:", ndSum, "stSum:", stSum)
	// 		if p == nums[i] {
	// 			// this is not the trionic subarray, reset
	// 			l, p = nums[i], nums[i]
	// 			stSum, ndSum = 0, 0
	// 			i = i - 1
	// 			continue
	// 		}
	// 		if p < nums[i] {
	// 			// which condition to reset p pos and val here?
	// 			if !pUpdated {
	// 				// q has not been set yet so this indicate a reset
	// 				p = math.MaxInt32
	// 				l = nums[i]
	// 				i = i - 1
	// 				lUpdated = false
	// 				continue
	// 			}
	//
	// 			// this maybe the third state
	// 			fmt.Println("Found potential third state starting at index", i, "with value", nums[i])
	// 			q = p
	// 			for i < end && q < nums[i] {
	// 				rdSum += int64(q)
	// 				q = nums[i]
	// 				qUpdated = true
	// 				i++
	// 			}
	// 			if i == end {
	// 				rdSum += int64(q)
	// 				break
	// 			}
	// 			fmt.Println("Exited third loop at index", i, " l:", l, " p: ", p, " q: ", q, "rdSum:", rdSum, "ndSum:", ndSum, "stSum:", stSum)
	// 			if q == nums[i] {
	// 				// t is not the trionic subarray, reset
	// 				l, p, q = nums[i], nums[i], nums[i]
	// 				stSum, ndSum, rdSum = 0, 0, 0
	// 				i = i - 1
	// 				continue
	// 			}
	// 			if q > nums[i] {
	// 				if !qUpdated {
	// 					// reset game
	// 					p = math.MaxInt32
	// 					q = math.MinInt32
	// 					l = nums[i]
	// 					i = i - 1
	// 					qUpdated = false
	// 					lUpdated = false
	// 					continue
	// 				} else {
	// 					rdSum += int64(q)
	// 					// probably the end, break
	// 					break
	// 				}
	// 			}
	// 		}
	// 	}
	// }
	//
	// fmt.Println("stSum:", stSum, "ndSum:", ndSum, "rdSum:", rdSum)
	// return stSum + ndSum + rdSum
