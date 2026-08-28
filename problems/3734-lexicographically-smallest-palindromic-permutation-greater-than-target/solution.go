package main

// 3734. Lexicographically Smallest Palindromic Permutation Greater Than Target
//
// A palindrome is fully determined by its first half (plus the middle character
// when n is odd), so instead of enumerating permutations of s we only choose the
// first half from the multiset cnt[c]/2. To get the smallest palindrome strictly
// greater than target we prefer the candidate that agrees with target on the
// longest prefix:
//
//  1. If target's own first half is exactly the available half multiset, the
//     palindrome built from it is the unique candidate matching target on the
//     first h characters. Return it when it is strictly greater than target.
//  2. Otherwise pick a break position i < h, as large as possible: keep
//     target[:i] in the first half, place the smallest still-available character
//     greater than target[i] at i, and dump everything left over in ascending
//     order. The string already exceeds target at i, so the sorted tail is the
//     smallest legal completion.
func smallestPalindrome(s string, target string) string {
	n := len(s)

	var cnt [26]int
	for i := 0; i < n; i++ {
		cnt[s[i]-'a']++
	}

	// A palindromic permutation exists only when at most one letter is odd.
	odd, oddCount := 0, 0
	for c := 0; c < 26; c++ {
		if cnt[c]%2 == 1 {
			odd, oddCount = c, oddCount+1
		}
	}
	if oddCount > 1 {
		return ""
	}

	h := n / 2
	var half [26]int
	for c := 0; c < 26; c++ {
		half[c] = cnt[c] / 2
	}

	build := func(firstHalf []byte) string {
		b := make([]byte, n)
		copy(b, firstHalf)
		if n%2 == 1 {
			b[h] = byte('a' + odd)
		}
		for i := 0; i < h; i++ {
			b[n-1-i] = firstHalf[i]
		}
		return string(b)
	}

	// Step 1: consume target[:h] greedily, recording the multiset left over
	// after every prefix length. states[i] is what remains after target[:i].
	states := make([][26]int, 1, h+1)
	states[0] = half
	cur := half
	maxPrefix := h
	for i := 0; i < h; i++ {
		c := int(target[i] - 'a')
		if cur[c] == 0 {
			maxPrefix = i
			break
		}
		cur[c]--
		states = append(states, cur)
	}

	// The whole first half can mirror target's first half: that is the single
	// best-prefixed candidate, so it wins whenever it beats target.
	if maxPrefix == h {
		if cand := build([]byte(target[:h])); cand > target {
			return cand
		}
	}

	// Step 2: break as late as possible inside the first half.
	start := maxPrefix
	if start > h-1 {
		start = h - 1
	}
	for i := start; i >= 0; i-- {
		rem := states[i]
		for c := int(target[i]-'a') + 1; c < 26; c++ {
			if rem[c] == 0 {
				continue
			}
			rem[c]--
			firstHalf := make([]byte, 0, h)
			firstHalf = append(firstHalf, target[:i]...)
			firstHalf = append(firstHalf, byte('a'+c))
			for d := 0; d < 26; d++ {
				for k := 0; k < rem[d]; k++ {
					firstHalf = append(firstHalf, byte('a'+d))
				}
			}
			return build(firstHalf)
		}
	}

	return ""
}
