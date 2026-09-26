package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name      string
		s         string
		knowledge [][]string
		expected  string
	}{
		{"example 1: two known keys", "(name)is(age)yearsold", [][]string{{"name", "bob"}, {"age", "two"}}, "bobistwoyearsold"},
		{"example 2: unknown key becomes ?", "hi(name)", [][]string{{"a", "b"}}, "hi?"},
		{"example 3: repeated key, plain letters untouched", "(a)(a)(a)aaa", [][]string{{"a", "yes"}}, "yesyesyesaaa"},
		{"edge case: no brackets", "hello", [][]string{{"hello", "x"}}, "hello"},
		{"edge case: empty knowledge", "(x)y(z)", [][]string{}, "?y?"},
		{"edge case: single bracket pair only", "(k)", [][]string{{"k", "value"}}, "value"},
		{"edge case: mixed known and unknown", "a(b)c(d)e", [][]string{{"d", "zz"}}, "a?czze"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := evaluate(tt.s, tt.knowledge)
			if result != tt.expected {
				t.Errorf("got %q, want %q", result, tt.expected)
			}
		})
	}
}
