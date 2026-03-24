package main

// Construct Product Matrix using prefix and suffix products modulo 12345.
// Flatten the 2D grid conceptually. Forward pass stores prefix products,
// backward pass multiplies in suffix products. O(n*m) time, O(1) extra space.
func constructProductMatrix(grid [][]int) [][]int {
	const mod = 12345
	n := len(grid)
	m := len(grid[0])

	// Initialize result matrix
	p := make([][]int, n)
	for i := range p {
		p[i] = make([]int, m)
	}

	// Forward pass: store prefix products
	prefix := 1
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p[i][j] = prefix
			prefix = prefix * (grid[i][j] % mod) % mod
		}
	}

	// Backward pass: multiply by suffix products
	suffix := 1
	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			p[i][j] = p[i][j] * suffix % mod
			suffix = suffix * (grid[i][j] % mod) % mod
		}
	}

	return p
}
