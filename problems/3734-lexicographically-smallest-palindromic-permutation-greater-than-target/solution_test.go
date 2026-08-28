package main

import (
	"math/rand"
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
		{"example 1: next palindrome after a valid one", "baba", "abba", "baab"},
		{"example 2: every palindromic permutation is too small", "baba", "bbaa", ""},
		{"example 3: no palindromic permutation exists", "abc", "abb", ""},
		{"example 4: the only palindrome already beats target", "aac", "abb", "aca"},

		{"edge case: n == 1, equal so not strictly greater", "a", "a", ""},
		{"edge case: n == 1, single letter beats target", "b", "a", "b"},
		{"edge case: target equals a palindromic permutation", "abba", "abba", "baab"},
		{"edge case: target prefix not buildable from half multiset", "aabb", "aaaa", "abba"},
		{"edge case: target above every palindromic permutation", "aabb", "zzzz", ""},
		{"edge case: odd length with forced middle character", "aabbc", "aaaaa", "abcba"},
		{"edge case: two odd counts on even length", "ab", "aa", ""},
		{"edge case: all identical letters, target smaller", "aaaa", "aaab", ""},
		{"edge case: all identical letters, target smaller everywhere", "bbbb", "aaaa", "bbbb"},
		{"edge case: break must happen at the very first position", "aabb", "abbb", "baab"},
		{"edge case: mirrored half is exactly target's half", "abab", "abab", "abba"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := smallestPalindrome(tt.s, tt.target)
			if got != tt.expected {
				t.Errorf("smallestPalindrome(%q, %q) = %q, want %q", tt.s, tt.target, got, tt.expected)
			}
		})
	}
}

// bruteForce enumerates every distinct permutation of s and returns the
// lexicographically smallest palindrome strictly greater than target.
func bruteForce(s, target string) string {
	var cnt [26]int
	for i := 0; i < len(s); i++ {
		cnt[s[i]-'a']++
	}

	n := len(s)
	best := ""
	buf := make([]byte, 0, n)

	var isPalindrome func(b []byte) bool
	isPalindrome = func(b []byte) bool {
		for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
			if b[i] != b[j] {
				return false
			}
		}
		return true
	}

	var rec func()
	rec = func() {
		if len(buf) == n {
			cand := string(buf)
			if isPalindrome(buf) && cand > target && (best == "" || cand < best) {
				best = cand
			}
			return
		}
		for c := 0; c < 26; c++ {
			if cnt[c] == 0 {
				continue
			}
			cnt[c]--
			buf = append(buf, byte('a'+c))
			rec()
			buf = buf[:len(buf)-1]
			cnt[c]++
		}
	}
	rec()

	return best
}

func TestSolutionAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(20260828))

	for iter := 0; iter < 400; iter++ {
		n := 1 + rng.Intn(6)
		alphabet := 1 + rng.Intn(3)

		letters := make([]byte, n)
		for i := range letters {
			letters[i] = byte('a' + rng.Intn(alphabet))
		}
		s := string(letters)

		// Half the time derive target from a shuffle of s so near-miss
		// prefixes get exercised, otherwise pick it uniformly at random.
		var target string
		if rng.Intn(2) == 0 {
			shuffled := append([]byte(nil), letters...)
			rng.Shuffle(n, func(i, j int) { shuffled[i], shuffled[j] = shuffled[j], shuffled[i] })
			target = string(shuffled)
		} else {
			raw := make([]byte, n)
			for i := range raw {
				raw[i] = byte('a' + rng.Intn(alphabet))
			}
			target = string(raw)
		}

		want := bruteForce(s, target)
		got := smallestPalindrome(s, target)
		if got != want {
			t.Fatalf("smallestPalindrome(%q, %q) = %q, want %q", s, target, got, want)
		}

		// The answer must be a permutation of s when one exists.
		if got != "" {
			a, b := []byte(got), []byte(s)
			sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
			sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
			if string(a) != string(b) {
				t.Fatalf("smallestPalindrome(%q, %q) = %q is not a permutation of s", s, target, got)
			}
		}
	}
}
