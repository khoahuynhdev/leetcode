package main

// 3518. Smallest Palindromic Rearrangement II
//
// A palindrome is completely determined by its first half plus an optional
// middle character, and comparing two palindromes of equal length
// lexicographically is the same as comparing their first halves. So the task
// reduces to: find the k-th lexicographically smallest distinct permutation of
// the multiset containing cnt[c]/2 copies of each character c.
//
// That half is built greedily, one position at a time. For each candidate
// character (in alphabetical order) we count the distinct permutations of the
// remaining multiset. If that count is >= k the answer starts with this
// character; otherwise we subtract the count and try the next character.
//
// The permutation counts are astronomically large, but k <= 1e6, so every count
// is clamped at 1e6. Both the multinomial product and the binomial coefficients
// inside it bail out as soon as they cross the cap, which keeps all arithmetic
// inside int64 and makes a single count O(26 * 25) in the worst case: a
// clamped binomial needs at most ~25 iterations because C(2i, i) passes 1e6 by
// i = 12.

// permLimit is the clamp for permutation counts. It is >= the maximum k, which
// is all the comparisons below need.
const permLimit = 1000000

func smallestPalindrome(s string, k int) string {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	var half [26]int
	mid := ""
	odds := 0
	for c := 0; c < 26; c++ {
		half[c] = cnt[c] / 2
		if cnt[c]%2 == 1 {
			odds++
			mid = string(rune('a' + c))
		}
	}
	// s is guaranteed palindromic, so at most one character has an odd count.
	// Guard anyway: more than one odd count admits no palindrome at all.
	if odds > 1 {
		return ""
	}

	if countPermutations(&half) < k {
		return ""
	}

	m := len(s) / 2
	first := make([]byte, 0, m)
	for len(first) < m {
		// Once k == 1 the smallest remaining arrangement is simply the
		// remaining characters in sorted order.
		if k == 1 {
			for c := 0; c < 26; c++ {
				for ; half[c] > 0; half[c]-- {
					first = append(first, byte('a'+c))
				}
			}
			break
		}
		for c := 0; c < 26; c++ {
			if half[c] == 0 {
				continue
			}
			half[c]--
			ways := countPermutations(&half)
			if ways >= k {
				first = append(first, byte('a'+c))
				break
			}
			// ways < k <= permLimit, so ways is an exact (unclamped) count.
			k -= ways
			half[c]++
		}
	}

	out := make([]byte, 0, len(s))
	out = append(out, first...)
	out = append(out, mid...)
	for i := len(first) - 1; i >= 0; i-- {
		out = append(out, first[i])
	}
	return string(out)
}

// countPermutations returns the number of distinct permutations of the multiset
// described by cnt, clamped to permLimit. It computes the multinomial
// coefficient as a product of binomials: placing the characters one group at a
// time, each group chooses its positions among the slots used so far.
func countPermutations(cnt *[26]int) int {
	res := int64(1)
	used := 0
	for c := 0; c < 26; c++ {
		if cnt[c] == 0 {
			continue
		}
		used += cnt[c]
		// res < permLimit and the binomial is <= permLimit, so the product
		// stays below 1e12 and cannot overflow.
		res *= int64(binomialCapped(used, cnt[c]))
		if res >= permLimit {
			return permLimit
		}
	}
	return int(res)
}

// binomialCapped returns C(n, r) clamped to permLimit. The running value after
// step i is exactly C(n-r+i, i), so every division is exact, and the value is
// non-decreasing in i, which makes the early exit safe.
func binomialCapped(n, r int) int {
	if r > n-r {
		r = n - r
	}
	res := int64(1)
	for i := 1; i <= r; i++ {
		// res < permLimit (1e6) and n <= 1e4, so res*(n-r+i) fits in int64.
		res = res * int64(n-r+i) / int64(i)
		if res >= permLimit {
			return permLimit
		}
	}
	return int(res)
}
