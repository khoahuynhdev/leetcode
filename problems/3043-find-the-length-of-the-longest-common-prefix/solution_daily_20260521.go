package main

// Approach: store every prefix of every number in arr1 in a hash set
// (a prefix is produced by repeated integer division by 10). Then for each
// number in arr2, strip its digits from the right (longest prefix first) and
// stop at the first prefix found in the set — that is the longest common
// prefix for that number. The answer is the max digit count over all of arr2.
// This avoids the O(N*M) pairwise comparison; it runs in O((N+M)*D) time
// where D is the number of digits (at most 9 for values <= 10^8).

func longestCommonPrefix(arr1 []int, arr2 []int) int {
	prefixes := make(map[int]struct{})
	for _, x := range arr1 {
		for x > 0 {
			prefixes[x] = struct{}{}
			x /= 10
		}
	}

	best := 0
	for _, y := range arr2 {
		length := digitCount(y)
		for y > 0 {
			if _, ok := prefixes[y]; ok {
				if length > best {
					best = length
				}
				break
			}
			y /= 10
			length--
		}
	}
	return best
}

func digitCount(n int) int {
	c := 0
	for n > 0 {
		c++
		n /= 10
	}
	return c
}
