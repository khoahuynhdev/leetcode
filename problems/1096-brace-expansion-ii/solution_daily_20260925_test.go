package main

import (
	"reflect"
	"testing"
)

func TestBraceExpansionII(t *testing.T) {
	tests := []struct {
		name       string
		expression string
		expected   []string
	}{
		{"example 1: concatenation of two brace groups", "{a,b}{c,{d,e}}", []string{"ac", "ad", "ae", "bc", "bd", "be"}},
		{"example 2: union with duplicates removed", "{{a,z},a{b,c},{ab,z}}", []string{"a", "ab", "ac", "z"}},
		{"description: simple union", "{a,b,c}", []string{"a", "b", "c"}},
		{"description: nested union with overlap", "{{a,b},{b,c}}", []string{"a", "b", "c"}},
		{"description: mixed letters and groups", "a{b,c}{d,e}f{g,h}", []string{"abdfg", "abdfh", "abefg", "abefh", "acdfg", "acdfh", "acefg", "acefh"}},
		{"edge case: single letter", "a", []string{"a"}},
		{"edge case: plain word without braces", "abc", []string{"abc"}},
		{"edge case: union where one branch concatenates", "{a,b}c{d,e}", []string{"acd", "ace", "bcd", "bce"}},
		{"edge case: products collapse to duplicates", "{a,ab}{c,bc}", []string{"abbc", "abc", "ac"}},
		{"edge case: deep nesting", "{{{a}}}", []string{"a"}},
		{"edge case: sorted output across different lengths", "{z,{y,x}}{b,a}", []string{"xa", "xb", "ya", "yb", "za", "zb"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := braceExpansionII(tt.expression)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("braceExpansionII(%q) = %v, want %v", tt.expression, result, tt.expected)
			}
		})
	}
}
