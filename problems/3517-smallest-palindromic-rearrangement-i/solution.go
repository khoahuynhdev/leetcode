package main

// 3517. Smallest Palindromic Rearrangement I
// https://leetcode.com/problems/smallest-palindromic-rearrangement-i/
//
// Approach: counting sort over the 26-letter alphabet.
//
// A palindrome is determined entirely by its left half plus an optional middle
// character, so comparing two palindromes of equal length reduces to comparing
// their left halves. The multiset of the left half is fixed: each letter c
// contributes cnt[c]/2 copies. The lexicographically smallest arrangement of a
// fixed multiset is its ascending order, so we sweep 'a'..'z' and emit
// cnt[c]/2 copies of each, writing every character and its mirror at once.
//
// Since s is guaranteed palindromic, at most one letter has an odd count; that
// leftover copy is forced into the center when the length is odd.
//
// Time: O(n + 26). Space: O(n) for the result, O(1) auxiliary.
func smallestPalindrome(s string) string {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	n := len(s)
	res := make([]byte, n)

	// Fill the left half in ascending order, mirroring each write to the right.
	left := 0
	for c := 0; c < 26; c++ {
		for k := cnt[c] / 2; k > 0; k-- {
			ch := byte('a' + c)
			res[left] = ch
			res[n-1-left] = ch
			left++
		}
	}

	// Odd length: exactly one letter has an odd count and must sit dead center.
	if n%2 == 1 {
		for c := 0; c < 26; c++ {
			if cnt[c]%2 == 1 {
				res[n/2] = byte('a' + c)
				break
			}
		}
	}

	return string(res)
}
