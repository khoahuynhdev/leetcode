package main

// Maximum Ice Cream Bars (LeetCode 1833)
//
// Approach: greedy + counting sort. Buying the cheapest bars first maximizes
// the number purchased. Since costs are bounded (1 <= costs[i] <= 10^5), we use
// counting sort: bucket bars by price into a frequency array, then walk prices
// from cheapest to most expensive, buying as many as the budget allows at each
// level. Runs in O(n + maxCost) time.
func maxIceCream(costs []int, coins int) int {
	if len(costs) == 0 {
		return 0
	}

	// Find the maximum cost to size the frequency table.
	maxCost := 0
	for _, c := range costs {
		if c > maxCost {
			maxCost = c
		}
	}

	// count[p] is the number of bars priced at p.
	count := make([]int, maxCost+1)
	for _, c := range costs {
		count[c]++
	}

	bought := 0
	for price := 1; price <= maxCost; price++ {
		if count[price] == 0 {
			continue
		}

		// How many bars at this price can we afford?
		affordable := coins / price
		if affordable == 0 {
			// Cannot afford even one at this price; all higher prices are
			// also unaffordable, so we can stop early.
			break
		}
		if affordable > count[price] {
			affordable = count[price]
		}

		bought += affordable
		coins -= affordable * price
	}

	return bought
}
