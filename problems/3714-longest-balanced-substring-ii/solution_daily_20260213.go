package main

// Approach: Split into cases by number of distinct characters in the balanced substring.
// 1) Single char: longest run of one character.
// 2) Two chars: for each pair, find maximal segments with only those chars,
//    then use prefix-difference technique within each segment.
// 3) All three chars: use 2D prefix-difference state (countA-countB, countA-countC).
// Time: O(n), Space: O(n)

func longestBalancedSubstring(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	ans := 1 // single character is always balanced

	// Case 1: longest run of a single character
	run := 1
	for i := 1; i < n; i++ {
		if s[i] == s[i-1] {
			run++
		} else {
			run = 1
		}
		if run > ans {
			ans = run
		}
	}

	// Case 2: two-character pairs
	pairs := [][2]byte{{'a', 'b'}, {'a', 'c'}, {'b', 'c'}}
	for _, pair := range pairs {
		c1, c2 := pair[0], pair[1]
		// Process maximal segments containing only c1 and c2
		i := 0
		for i < n {
			if s[i] != c1 && s[i] != c2 {
				i++
				continue
			}
			// Start of a segment with only c1/c2
			diff := 0 // count(c1) - count(c2)
			first := map[int]int{0: i - 1}
			j := i
			for j < n && (s[j] == c1 || s[j] == c2) {
				if s[j] == c1 {
					diff++
				} else {
					diff--
				}
				if idx, ok := first[diff]; ok {
					length := j - idx
					if length > ans {
						ans = length
					}
				} else {
					first[diff] = j
				}
				j++
			}
			i = j
		}
	}

	// Case 3: all three characters with equal counts
	// State: (countA - countB, countA - countC)
	type state struct {
		ab, ac int
	}
	firstSeen := map[state]int{{0, 0}: -1}
	var ca, cb, cc int
	for i := 0; i < n; i++ {
		switch s[i] {
		case 'a':
			ca++
		case 'b':
			cb++
		case 'c':
			cc++
		}
		st := state{ca - cb, ca - cc}
		if idx, ok := firstSeen[st]; ok {
			length := i - idx
			if length > ans {
				ans = length
			}
		} else {
			firstSeen[st] = i
		}
	}

	return ans
}
