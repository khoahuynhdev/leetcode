package main

import (
	"math/rand"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name            string
		s               string
		queryCharacters string
		queryIndices    []int
		expected        []int
	}{
		{
			name:            "example 1: babacc with three overlapping updates",
			s:               "babacc",
			queryCharacters: "bcb",
			queryIndices:    []int{1, 3, 3},
			expected:        []int{3, 3, 4},
		},
		{
			name:            "example 2: abyzz where the answer grows",
			s:               "abyzz",
			queryCharacters: "aa",
			queryIndices:    []int{2, 1},
			expected:        []int{2, 3},
		},
		{
			name:            "edge case: single character string, update is always length 1",
			s:               "a",
			queryCharacters: "bz",
			queryIndices:    []int{0, 0},
			expected:        []int{1, 1},
		},
		{
			name:            "edge case: update splits a long run so the answer drops",
			s:               "aaaaa",
			queryCharacters: "b",
			queryIndices:    []int{2},
			expected:        []int{2},
		},
		{
			name:            "edge case: update joins two runs across a segment boundary",
			s:               "aabaa",
			queryCharacters: "a",
			queryIndices:    []int{2},
			expected:        []int{5},
		},
		{
			name:            "edge case: no-op query writing the character already present",
			s:               "zzzz",
			queryCharacters: "z",
			queryIndices:    []int{3},
			expected:        []int{4},
		},
		{
			name:            "edge case: whole string collapses to one repeating character",
			s:               "abc",
			queryCharacters: "cc",
			queryIndices:    []int{0, 1},
			expected:        []int{1, 3},
		},
		{
			name:            "edge case: answers rise and fall across successive queries",
			s:               "abab",
			queryCharacters: "aba",
			queryIndices:    []int{1, 1, 3},
			expected:        []int{3, 1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestRepeating(tt.s, tt.queryCharacters, tt.queryIndices)
			if len(got) != len(tt.expected) {
				t.Fatalf("got %v (len %d), want %v (len %d)", got, len(got), tt.expected, len(tt.expected))
			}
			for i := range got {
				if got[i] != tt.expected[i] {
					t.Fatalf("got %v, want %v (first mismatch at index %d)", got, tt.expected, i)
				}
			}
		})
	}
}

// bruteForceLongestRepeating rescans the whole string after every query. It is
// O(n*k) and far too slow for the real constraints, but it is obviously correct
// and makes a good oracle for randomised cross-checking.
func bruteForceLongestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	b := []byte(s)
	lengths := make([]int, len(queryIndices))
	for i, pos := range queryIndices {
		b[pos] = queryCharacters[i]

		best, run := 1, 1
		for j := 1; j < len(b); j++ {
			if b[j] == b[j-1] {
				run++
			} else {
				run = 1
			}
			if run > best {
				best = run
			}
		}
		lengths[i] = best
	}
	return lengths
}

func TestSolutionAgainstBruteForce(t *testing.T) {
	rng := rand.New(rand.NewSource(2213))

	// A small alphabet makes long runs (and boundary merges) frequent.
	const alphabet = "abc"

	for trial := 0; trial < 200; trial++ {
		n := 1 + rng.Intn(20)
		k := 1 + rng.Intn(20)

		s := make([]byte, n)
		for i := range s {
			s[i] = alphabet[rng.Intn(len(alphabet))]
		}
		qc := make([]byte, k)
		qi := make([]int, k)
		for i := 0; i < k; i++ {
			qc[i] = alphabet[rng.Intn(len(alphabet))]
			qi[i] = rng.Intn(n)
		}

		got := longestRepeating(string(s), string(qc), qi)
		want := bruteForceLongestRepeating(string(s), string(qc), qi)
		for i := range want {
			if got[i] != want[i] {
				t.Fatalf("s=%q queryCharacters=%q queryIndices=%v: got %v, want %v",
					s, qc, qi, got, want)
			}
		}
	}
}
