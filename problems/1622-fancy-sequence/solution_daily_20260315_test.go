package main

import "testing"

func TestFancy(t *testing.T) {
	type op struct {
		method string
		arg    int
		want   int // only used for getIndex
	}

	tests := []struct {
		name string
		ops  []op
	}{
		{
			name: "example 1: append, addAll, multAll, getIndex",
			ops: []op{
				{"append", 2, 0},
				{"addAll", 3, 0},
				{"append", 7, 0},
				{"multAll", 2, 0},
				{"getIndex", 0, 10},
				{"addAll", 3, 0},
				{"append", 10, 0},
				{"multAll", 2, 0},
				{"getIndex", 0, 26},
				{"getIndex", 1, 34},
				{"getIndex", 2, 20},
			},
		},
		{
			name: "edge case: getIndex on empty sequence",
			ops: []op{
				{"getIndex", 0, -1},
				{"getIndex", 5, -1},
			},
		},
		{
			name: "edge case: index out of bounds after appends",
			ops: []op{
				{"append", 1, 0},
				{"append", 2, 0},
				{"getIndex", 0, 1},
				{"getIndex", 1, 2},
				{"getIndex", 2, -1},
				{"getIndex", 100, -1},
			},
		},
		{
			name: "edge case: only addAll and multAll without affecting new append",
			ops: []op{
				{"append", 5, 0},
				{"addAll", 10, 0},
				{"multAll", 3, 0},
				// (5+10)*3 = 45
				{"getIndex", 0, 45},
				// append after transforms; new value is unaffected by prior ops
				{"append", 7, 0},
				{"getIndex", 1, 7},
				// but prior element still has transforms
				{"getIndex", 0, 45},
			},
		},
		{
			name: "edge case: multiple multiplications and additions",
			ops: []op{
				{"append", 1, 0},
				{"multAll", 100, 0},
				{"addAll", 100, 0},
				{"multAll", 100, 0},
				{"addAll", 100, 0},
				// ((1*100)+100)*100+100 = (200)*100+100 = 20100
				{"getIndex", 0, 20100},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fancy := Constructor()
			for i, o := range tt.ops {
				switch o.method {
				case "append":
					fancy.Append(o.arg)
				case "addAll":
					fancy.AddAll(o.arg)
				case "multAll":
					fancy.MultAll(o.arg)
				case "getIndex":
					got := fancy.GetIndex(o.arg)
					if got != o.want {
						t.Errorf("op %d: GetIndex(%d) = %d, want %d", i, o.arg, got, o.want)
					}
				}
			}
		})
	}
}
