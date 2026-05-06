package main

// rotateTheBox simulates gravity row-by-row in the original orientation
// (since 90° clockwise rotation maps "right" to "down"), then rotates the
// result 90° clockwise.
//
// For each row, a right-to-left two-pointer scan tracks the rightmost empty
// slot a stone can fall into. Obstacles reset that slot to their left.
func rotateTheBox(boxGrid [][]byte) [][]byte {
	m := len(boxGrid)
	n := len(boxGrid[0])

	for i := 0; i < m; i++ {
		empty := n - 1
		for j := n - 1; j >= 0; j-- {
			switch boxGrid[i][j] {
			case '*':
				empty = j - 1
			case '#':
				boxGrid[i][j] = '.'
				boxGrid[i][empty] = '#'
				empty--
			}
		}
	}

	result := make([][]byte, n)
	for i := 0; i < n; i++ {
		result[i] = make([]byte, m)
		for j := 0; j < m; j++ {
			result[i][j] = boxGrid[m-1-j][i]
		}
	}
	return result
}
