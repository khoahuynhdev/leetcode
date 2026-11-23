package main

import "testing"

func TestSolution(t *testing.T) {
	// Each op is either "add" (arg = num) or "get" (arg = k, want = expected).
	type op struct {
		kind string
		arg  int
		want int
	}

	tests := []struct {
		name string
		ops  []op
	}{
		{
			name: "example 1: stream with a zero in the middle",
			ops: []op{
				{"add", 3, 0}, {"add", 0, 0}, {"add", 2, 0}, {"add", 5, 0}, {"add", 4, 0},
				{"get", 2, 20}, {"get", 3, 40}, {"get", 4, 0},
				{"add", 8, 0},
				{"get", 2, 32},
			},
		},
		{
			name: "edge case: no zeros, k equals whole stream",
			ops: []op{
				{"add", 2, 0}, {"add", 3, 0}, {"add", 4, 0},
				{"get", 3, 24}, {"get", 1, 4},
			},
		},
		{
			name: "edge case: zero is the most recent element",
			ops: []op{
				{"add", 5, 0}, {"add", 0, 0},
				{"get", 1, 0}, {"get", 2, 0},
			},
		},
		{
			name: "edge case: zero at the start, k reaches exactly the zero",
			ops: []op{
				{"add", 0, 0}, {"add", 2, 0},
				{"get", 1, 2}, {"get", 2, 0},
			},
		},
		{
			name: "edge case: multiple zeros and ones",
			ops: []op{
				{"add", 0, 0}, {"add", 7, 0}, {"add", 0, 0}, {"add", 1, 0}, {"add", 1, 0}, {"add", 9, 0},
				{"get", 3, 9}, {"get", 4, 0}, {"get", 6, 0},
			},
		},
		{
			name: "edge case: single element",
			ops: []op{
				{"add", 100, 0},
				{"get", 1, 100},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Constructor()
			for i, o := range tt.ops {
				switch o.kind {
				case "add":
					p.Add(o.arg)
				case "get":
					if got := p.GetProduct(o.arg); got != o.want {
						t.Errorf("op %d: GetProduct(%d) = %d, want %d", i, o.arg, got, o.want)
					}
				}
			}
		})
	}
}
