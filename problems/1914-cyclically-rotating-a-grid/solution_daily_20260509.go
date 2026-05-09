package main

// Rotate each concentric layer independently.
// For each layer, flatten the perimeter clockwise into a slice, left-shift it
// by (k mod perimeter length) — which is the same as counter-clockwise rotation
// when the buffer is laid out clockwise — then write it back along the same
// clockwise path.
func rotateGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	if m == 0 {
		return grid
	}
	n := len(grid[0])

	layers := m / 2
	if n/2 < layers {
		layers = n / 2
	}

	for layer := 0; layer < layers; layer++ {
		top, left := layer, layer
		bottom, right := m-1-layer, n-1-layer

		size := 2*((bottom-top)+(right-left)) // perimeter length

		ring := make([]int, 0, size)
		// top edge: left → right
		for j := left; j <= right; j++ {
			ring = append(ring, grid[top][j])
		}
		// right edge: top+1 → bottom
		for i := top + 1; i <= bottom; i++ {
			ring = append(ring, grid[i][right])
		}
		// bottom edge: right-1 → left
		for j := right - 1; j >= left; j-- {
			ring = append(ring, grid[bottom][j])
		}
		// left edge: bottom-1 → top+1
		for i := bottom - 1; i > top; i-- {
			ring = append(ring, grid[i][left])
		}

		shift := k % size

		// Write back in the same clockwise order, sourcing from rotated index.
		idx := 0
		for j := left; j <= right; j++ {
			grid[top][j] = ring[(idx+shift)%size]
			idx++
		}
		for i := top + 1; i <= bottom; i++ {
			grid[i][right] = ring[(idx+shift)%size]
			idx++
		}
		for j := right - 1; j >= left; j-- {
			grid[bottom][j] = ring[(idx+shift)%size]
			idx++
		}
		for i := bottom - 1; i > top; i-- {
			grid[i][left] = ring[(idx+shift)%size]
			idx++
		}
	}

	return grid
}
