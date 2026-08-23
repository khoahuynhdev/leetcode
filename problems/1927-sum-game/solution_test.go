package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		num      string
		expected bool
	}{
		{"example 1: no moves left, sums already equal", "5023", false},
		{"example 2: two '?' on the right, gap is not 9", "25??", true},
		{"example 3: gap of 9 with two unmatched '?'", "?3295???", false},

		{"edge case: minimum length, no '?', sums differ", "19", true},
		{"edge case: minimum length, no '?', sums equal", "77", false},
		{"edge case: minimum length, single '?' so Alice moves last", "?0", true},
		{"edge case: single '?' in the right half", "9?", true},

		{"edge case: one '?' per half cancels, known sums equal", "??", false},
		{"edge case: one '?' per half cancels, known gap of 9", "0?9?", true},
		{"edge case: three '?' per half, all matched", "??????", false},

		{"edge case: unmatched '?' on the left, gap is exactly 9", "??09", false},
		{"edge case: unmatched '?' on the left, gap is not 9", "??00", true},

		{"edge case: no '?', sums differ", "1234", true},
		{"edge case: all zeros", "000000", false},
		{"edge case: four unmatched '?' needing a gap of 18", "9540????", false},
		{"edge case: four unmatched '?' with a gap of 19", "9541????", true},
		{"edge case: '?' in both halves, only two unmatched", "99??????", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumGame(tt.num); got != tt.expected {
				t.Errorf("sumGame(%q) = %v, want %v", tt.num, got, tt.expected)
			}
		})
	}
}

// TestSolutionAgainstBruteForce cross-checks the O(n) rule against an exhaustive
// minimax search on every short board built from the digits 0, 1, 9 and '?'.
func TestSolutionAgainstBruteForce(t *testing.T) {
	alphabet := []byte{'0', '1', '9', '?'}

	for _, n := range []int{2, 4} {
		total := 1
		for i := 0; i < n; i++ {
			total *= len(alphabet)
		}

		for code := 0; code < total; code++ {
			num := make([]byte, n)
			for i, rest := 0, code; i < n; i++ {
				num[i] = alphabet[rest%len(alphabet)]
				rest /= len(alphabet)
			}

			want := bruteForceAliceWins(num, true)
			if got := sumGame(string(num)); got != want {
				t.Errorf("sumGame(%q) = %v, brute force says %v", num, got, want)
			}
		}
	}
}

// bruteForceAliceWins plays out every legal move sequence and returns whether Alice
// wins under optimal play. Only usable on very short inputs.
func bruteForceAliceWins(num []byte, aliceTurn bool) bool {
	remaining := false
	for _, c := range num {
		if c == '?' {
			remaining = true
			break
		}
	}

	if !remaining {
		half, sum := len(num)/2, 0
		for i, c := range num {
			if i < half {
				sum += int(c - '0')
			} else {
				sum -= int(c - '0')
			}
		}
		return sum != 0
	}

	for i := range num {
		if num[i] != '?' {
			continue
		}
		for d := byte('0'); d <= '9'; d++ {
			num[i] = d
			result := bruteForceAliceWins(num, !aliceTurn)
			num[i] = '?'

			// Alice takes any line that wins for her; Bob takes any line that does not.
			if result == aliceTurn {
				return aliceTurn
			}
		}
	}

	// Every move leads to the opponent's preferred outcome.
	return !aliceTurn
}
