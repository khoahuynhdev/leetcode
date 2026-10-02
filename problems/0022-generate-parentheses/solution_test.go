package main

import (
	"sort"
	"testing"
)

func isWellFormed(s string) bool {
	balance := 0
	for _, c := range s {
		if c == '(' {
			balance++
		} else {
			balance--
		}
		if balance < 0 {
			return false
		}
	}
	return balance == 0
}

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []string
	}{
		{"example 1: n = 3", 3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{"example 2: n = 1", 1, []string{"()"}},
		{"edge case: n = 2", 2, []string{"(())", "()()"}},
		{"edge case: n = 4", 4, []string{
			"(((())))", "((()()))", "((())())", "((()))()", "(()(()))",
			"(()()())", "(()())()", "(())(())", "(())()()", "()((()))",
			"()(()())", "()(())()", "()()(())", "()()()()",
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := generateParenthesis(tt.n)
			got := append([]string(nil), result...)
			want := append([]string(nil), tt.expected...)
			sort.Strings(got)
			sort.Strings(want)
			if len(got) != len(want) {
				t.Fatalf("got %d results %v, want %d results %v", len(got), got, len(want), want)
			}
			for i := range got {
				if got[i] != want[i] {
					t.Errorf("got %v, want %v", got, want)
					break
				}
			}
		})
	}
}

func TestGenerateParenthesisCatalanCount(t *testing.T) {
	catalan := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430}
	for n := 1; n <= 8; n++ {
		result := generateParenthesis(n)
		if len(result) != catalan[n] {
			t.Errorf("n=%d: got %d results, want %d", n, len(result), catalan[n])
		}
		seen := make(map[string]bool)
		for _, s := range result {
			if len(s) != 2*n || !isWellFormed(s) {
				t.Errorf("n=%d: invalid string %q", n, s)
			}
			if seen[s] {
				t.Errorf("n=%d: duplicate string %q", n, s)
			}
			seen[s] = true
		}
	}
}
