package main

import (
	"sort"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		target   string
		expected string
	}{
		{"example 1: breaks at index 1", "abc", "bba", "bca"},
		{"example 2: target[0] absent from s", "leet", "code", "eelt"},
		{"example 3: no permutation is greater", "baba", "bbaa", ""},
		{"edge case: single char, s greater", "b", "a", "b"},
		{"edge case: single char, equal", "a", "a", ""},
		{"edge case: single char, s smaller", "a", "b", ""},
		{"edge case: s is an anagram of target", "abc", "abc", "acb"},
		{"edge case: all identical letters, no answer", "aaa", "aab", ""},
		{"edge case: all identical letters, answer is s", "bbb", "aaa", "bbb"},
		{"edge case: full prefix match, breaks at last index", "aab", "aaa", "aab"},
		{"edge case: two letters, must break at index 0", "ba", "ab", "ba"},
		{"edge case: target far above every permutation", "abc", "zzz", ""},
		{"edge case: target far below every permutation", "zzz", "aaa", "zzz"},
		{"edge case: duplicates in s with long shared prefix", "aabbc", "aabbb", "aabbc"},
		{"edge case: backtracks past an unusable position", "abcd", "abdd", "acbd"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallestPermutation(tt.s, tt.target); got != tt.expected {
				t.Errorf("smallestPermutation(%q, %q) = %q, want %q", tt.s, tt.target, got, tt.expected)
			}
		})
	}
}

// TestSolutionAgainstBruteForce cross-checks the greedy against exhaustive
// permutation enumeration over every length-3 string pair on the alphabet
// {a, b, c} (729 pairs), which covers duplicates, ties and impossible cases.
func TestSolutionAgainstBruteForce(t *testing.T) {
	alphabet := []byte{'a', 'b', 'c'}
	var words []string
	for _, x := range alphabet {
		for _, y := range alphabet {
			for _, z := range alphabet {
				words = append(words, string([]byte{x, y, z}))
			}
		}
	}

	for _, s := range words {
		for _, target := range words {
			got := smallestPermutation(s, target)
			want := bruteForceSmallestPermutation(s, target)
			if got != want {
				t.Errorf("smallestPermutation(%q, %q) = %q, want %q", s, target, got, want)
			}
		}
	}
}

// bruteForceSmallestPermutation enumerates every permutation of s and returns
// the smallest one strictly greater than target, or "" if there is none.
func bruteForceSmallestPermutation(s, target string) string {
	best := ""
	chars := []byte(s)
	sort.Slice(chars, func(i, j int) bool { return chars[i] < chars[j] })

	var permute func(prefix []byte, used []bool)
	permute = func(prefix []byte, used []bool) {
		if len(prefix) == len(chars) {
			candidate := string(prefix)
			if candidate > target && (best == "" || candidate < best) {
				best = candidate
			}
			return
		}
		for i := range chars {
			if used[i] {
				continue
			}
			used[i] = true
			permute(append(prefix, chars[i]), used)
			used[i] = false
		}
	}

	permute(make([]byte, 0, len(chars)), make([]bool, len(chars)))
	return best
}
