// 3302. Find the Lexicographically Smallest Valid Sequence
//
// Greedy forward scan guarded by a precomputed suffix table.
//
// suffixMatch[i] is the length of the longest suffix of word2 that is a
// subsequence of word1[i:]. It is built right-to-left in O(n) with the ordinary
// subsequence-matching greedy run backwards, so suffixMatch[i] >= t answers
// "do the last t characters of word2 still fit exactly into word1[i:]?" in O(1).
//
// Scanning left-to-right we always take the smallest index that keeps a
// completion reachable:
//   - word1[i] == word2[j]: take i. A free match never costs the one allowed
//     change, so it dominates every alternative.
//   - otherwise, if the change is unspent and suffixMatch[i+1] >= m-1-j: spend
//     it on i. The guard proves the remainder can be matched exactly afterwards.
//   - otherwise i is unusable for word2[j], so skip it.
//
// Every appended index is the smallest one that preserves feasibility, which is
// exactly the lexicographically smallest array.
package main

func validSequence(word1 string, word2 string) []int {
	n, m := len(word1), len(word2)
	if m > n {
		return []int{}
	}

	// suffixMatch has a 0 sentinel at index n so the guard can read i+1 safely.
	suffixMatch := make([]int, n+1)
	for i := n - 1; i >= 0; i-- {
		k := suffixMatch[i+1]
		if k < m && word1[i] == word2[m-1-k] {
			k++
		}
		suffixMatch[i] = k
	}

	seq := make([]int, 0, m)
	changed := false
	j := 0
	for i := 0; i < n && j < m; i++ {
		switch {
		case word1[i] == word2[j]:
			seq = append(seq, i)
			j++
		case !changed && suffixMatch[i+1] >= m-1-j:
			seq = append(seq, i)
			changed = true
			j++
		}
	}

	if j < m {
		return []int{}
	}
	return seq
}
