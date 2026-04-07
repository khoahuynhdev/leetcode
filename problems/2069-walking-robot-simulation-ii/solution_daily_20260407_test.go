package main

import "testing"

func TestRobot(t *testing.T) {
	tests := []struct {
		name       string
		width      int
		height     int
		operations []string
		stepArgs   []int
		wantPos    [][]int    // nil entries for non-getPos operations
		wantDir    []*string  // nil entries for non-getDir operations
	}{
		{
			name:       "example 1: 6x3 grid with mixed operations",
			width:      6,
			height:     3,
			operations: []string{"step", "step", "getPos", "getDir", "step", "step", "step", "getPos", "getDir"},
			stepArgs:   []int{2, 2, 0, 0, 2, 1, 4, 0, 0},
			wantPos:    [][]int{nil, nil, {4, 0}, nil, nil, nil, nil, {1, 2}, nil},
			wantDir:    []*string{nil, nil, nil, strPtr("East"), nil, nil, nil, nil, strPtr("West")},
		},
		{
			name:       "edge case: no steps, initial position",
			width:      3,
			height:     3,
			operations: []string{"getPos", "getDir"},
			stepArgs:   []int{0, 0},
			wantPos:    [][]int{{0, 0}, nil},
			wantDir:    []*string{nil, strPtr("East")},
		},
		{
			name:       "edge case: full perimeter returns to origin facing South",
			width:      4,
			height:     3,
			operations: []string{"step", "getPos", "getDir"},
			stepArgs:   []int{10, 0, 0}, // P = 2*3 + 2*2 = 10
			wantPos:    [][]int{nil, {0, 0}, nil},
			wantDir:    []*string{nil, nil, strPtr("South")},
		},
		{
			name:       "edge case: minimum 2x2 grid",
			width:      2,
			height:     2,
			operations: []string{"step", "getPos", "getDir", "step", "getPos", "getDir", "step", "getPos", "getDir", "step", "getPos", "getDir"},
			stepArgs:   []int{1, 0, 0, 1, 0, 0, 1, 0, 0, 1, 0, 0},
			wantPos:    [][]int{nil, {1, 0}, nil, nil, {1, 1}, nil, nil, {0, 1}, nil, nil, {0, 0}, nil},
			wantDir:    []*string{nil, nil, strPtr("East"), nil, nil, strPtr("North"), nil, nil, strPtr("West"), nil, nil, strPtr("South")},
		},
		{
			name:       "edge case: large steps with modulo",
			width:      6,
			height:     3,
			operations: []string{"step", "getPos", "getDir"},
			stepArgs:   []int{100000, 0, 0},
			// P = 2*5 + 2*2 = 14, 100000 % 14 = 12 -> (0, 2), West (top-left corner)
			wantPos: [][]int{nil, {0, 2}, nil},
			wantDir: []*string{nil, nil, strPtr("West")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			robot := Constructor(tt.width, tt.height)
			for i, op := range tt.operations {
				switch op {
				case "step":
					robot.Step(tt.stepArgs[i])
				case "getPos":
					got := robot.GetPos()
					want := tt.wantPos[i]
					if want != nil && (got[0] != want[0] || got[1] != want[1]) {
						t.Errorf("operation %d getPos() = %v, want %v", i, got, want)
					}
				case "getDir":
					got := robot.GetDir()
					want := tt.wantDir[i]
					if want != nil && got != *want {
						t.Errorf("operation %d getDir() = %q, want %q", i, got, *want)
					}
				}
			}
		})
	}
}

func TestRobotCornerDirections(t *testing.T) {
	// Verify direction at each corner of a 3x3 grid (P=8)
	robot := Constructor(3, 3)

	// Walk to each corner and check direction
	corners := []struct {
		steps   int
		wantPos []int
		wantDir string
	}{
		{2, []int{2, 0}, "East"},   // bottom-right corner
		{2, []int{2, 2}, "North"},  // top-right corner
		{2, []int{0, 2}, "West"},   // top-left corner
		{2, []int{0, 0}, "South"},  // back to origin
	}

	for i, c := range corners {
		robot.Step(c.steps)
		pos := robot.GetPos()
		dir := robot.GetDir()
		if pos[0] != c.wantPos[0] || pos[1] != c.wantPos[1] {
			t.Errorf("corner %d: getPos() = %v, want %v", i, pos, c.wantPos)
		}
		if dir != c.wantDir {
			t.Errorf("corner %d: getDir() = %q, want %q", i, dir, c.wantDir)
		}
	}
}

func TestRobotMultipleLoops(t *testing.T) {
	// 5x4 grid, P = 2*4 + 2*3 = 14
	robot := Constructor(5, 4)
	// 3 full loops + 5 extra steps
	robot.Step(3*14 + 5)
	pos := robot.GetPos()
	dir := robot.GetDir()
	// pos 5 is on right edge: (4, 1), North
	if pos[0] != 4 || pos[1] != 1 {
		t.Errorf("getPos() = %v, want [4 1]", pos)
	}
	if dir != "North" {
		t.Errorf("getDir() = %q, want %q", dir, "North")
	}
}

func strPtr(s string) *string {
	return &s
}
