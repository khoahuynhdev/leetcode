package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestRemoveInvalidParentheses(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{"example 1: extra closing paren", "()())()", []string{"(())()", "()()()"}},
		{"example 2: letters mixed in", "(a)())()", []string{"(a())()", "(a)()()"}},
		{"example 3: fully invalid becomes empty", ")(", []string{""}},
		{"edge case: already valid", "(a)(b)", []string{"(a)(b)"}},
		{"edge case: letters only", "abc", []string{"abc"}},
		{"edge case: single open paren", "(", []string{""}},
		{"edge case: single close paren keeps letter", "x)", []string{"x"}},
		{"edge case: extra open parens", "((()", []string{"()"}},
		{"edge case: both unmatched sides", ")a(", []string{"a"}},
		{"edge case: nested with leading close", ")(()", []string{"()"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := removeInvalidParentheses(tt.input)
			sort.Strings(result)
			expected := append([]string(nil), tt.expected...)
			sort.Strings(expected)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("removeInvalidParentheses(%q) = %q, want %q", tt.input, result, expected)
			}
		})
	}
}
