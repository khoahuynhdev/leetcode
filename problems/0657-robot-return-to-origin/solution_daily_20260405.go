package main

// judgeCircle checks whether a robot following the given move sequence
// returns to the origin (0,0). It tracks net displacement on both axes
// and returns true only if both are zero after all moves.
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for _, m := range moves {
		switch m {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}
	}
	return x == 0 && y == 0
}
