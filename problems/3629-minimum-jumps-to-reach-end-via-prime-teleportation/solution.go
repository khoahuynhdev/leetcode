package main

// minJumps returns the minimum number of jumps required to walk from index 0
// to index n-1 of nums, where each jump is either an adjacent step (i±1) or,
// when nums[i] is a prime p, a teleport to any j with nums[j] % p == 0.
//
// Approach: BFS on indices. To avoid wiring up the O(n^2) teleport edges
// explicitly, we precompute, for each prime p that appears as some nums[i],
// the list of indices j with p | nums[j]. During BFS, the first time we visit
// an index whose value is the prime p, we enqueue that whole list and mark p
// as "expanded" so its bucket is never iterated again. This keeps the total
// work proportional to the sum of distinct prime factors over all elements.
func minJumps(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 0
	}

	maxVal := 0
	for _, v := range nums {
		if v > maxVal {
			maxVal = v
		}
	}

	// Smallest prime factor sieve up to maxVal.
	// spf[v] = smallest prime factor of v; spf[v] == v iff v is prime; spf[1] = 0.
	spf := make([]int, maxVal+1)
	for i := 2; i <= maxVal; i++ {
		if spf[i] == 0 {
			for j := i; j <= maxVal; j += i {
				if spf[j] == 0 {
					spf[j] = i
				}
			}
		}
	}

	isPrime := func(v int) bool {
		return v >= 2 && spf[v] == v
	}

	// Only primes that actually occur as a value in nums can ever be a teleport source.
	primeValues := make(map[int]bool)
	for _, v := range nums {
		if isPrime(v) {
			primeValues[v] = true
		}
	}

	// For each relevant prime p, collect the indices j whose nums[j] is divisible by p.
	primeToIdx := make(map[int][]int)
	for j, v := range nums {
		x := v
		for x > 1 {
			p := spf[x]
			if primeValues[p] {
				primeToIdx[p] = append(primeToIdx[p], j)
			}
			for x%p == 0 {
				x /= p
			}
		}
	}

	visited := make([]bool, n)
	primeExpanded := make(map[int]bool)
	visited[0] = true
	queue := []int{0}
	jumps := 0

	for len(queue) > 0 {
		next := make([]int, 0, len(queue))
		for _, idx := range queue {
			if idx == n-1 {
				return jumps
			}
			if idx > 0 && !visited[idx-1] {
				visited[idx-1] = true
				next = append(next, idx-1)
			}
			if idx+1 < n && !visited[idx+1] {
				visited[idx+1] = true
				next = append(next, idx+1)
			}
			v := nums[idx]
			if isPrime(v) && !primeExpanded[v] {
				primeExpanded[v] = true
				for _, j := range primeToIdx[v] {
					if !visited[j] {
						visited[j] = true
						next = append(next, j)
					}
				}
			}
		}
		queue = next
		jumps++
	}

	return -1
}
