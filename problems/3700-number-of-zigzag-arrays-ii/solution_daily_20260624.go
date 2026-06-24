package main

// Number of ZigZag Arrays II
//
// A ZigZag array must strictly alternate direction (up, down, up, down, ...),
// so the only thing that matters going forward is the last value and the last
// step's direction. Over m = r-l+1 relabeled values 1..m, track two states:
//   up[v]   = #valid prefixes ending at v whose last step went up   (prev < v)
//   down[v] = #valid prefixes ending at v whose last step went down (prev > v)
// Transition to a new last value w:
//   new_up[w]   = sum_{v < w} down[v]
//   new_down[w] = sum_{v > w} up[v]
// Base case (length 2): up[w] = w-1, down[w] = m-w.
//
// This is a fixed linear map on a vector of size 2m, applied n-2 times. Since
// n can be up to 1e9, raise the 2m x 2m transition matrix to the power n-2 with
// fast matrix exponentiation, multiply by the base vector, and sum all entries.
//
// Time:  O((2m)^3 * log n)   Space: O((2m)^2)

const mod = 1_000_000_007

type matrix [][]int64

func newMatrix(rows, cols int) matrix {
	m := make(matrix, rows)
	for i := range m {
		m[i] = make([]int64, cols)
	}
	return m
}

func multiply(a, b matrix) matrix {
	rows, inner, cols := len(a), len(b), len(b[0])
	c := newMatrix(rows, cols)
	for i := 0; i < rows; i++ {
		ai, ci := a[i], c[i]
		for k := 0; k < inner; k++ {
			if ai[k] == 0 {
				continue
			}
			v, bk := ai[k], b[k]
			for j := 0; j < cols; j++ {
				ci[j] = (ci[j] + v*bk[j]) % mod
			}
		}
	}
	return c
}

func matpow(m matrix, p int) matrix {
	n := len(m)
	res := newMatrix(n, n)
	for i := 0; i < n; i++ {
		res[i][i] = 1
	}
	for p > 0 {
		if p&1 == 1 {
			res = multiply(res, m)
		}
		m = multiply(m, m)
		p >>= 1
	}
	return res
}

func zigZagArrays(n int, l int, r int) int {
	m := r - l + 1
	size := 2 * m
	// Layout: index v-1 holds up[v]; index m+v-1 holds down[v], for v in 1..m.

	// Base vector for length-2 arrays.
	base := newMatrix(size, 1)
	for v := 1; v <= m; v++ {
		base[v-1][0] = int64(v - 1) // up[v]
		base[m+v-1][0] = int64(m - v) // down[v]
	}

	// Transition matrix.
	T := newMatrix(size, size)
	for w := 1; w <= m; w++ {
		for v := 1; v < w; v++ { // new_up[w] += down[v]
			T[w-1][m+v-1] = 1
		}
		for v := w + 1; v <= m; v++ { // new_down[w] += up[v]
			T[m+w-1][v-1] = 1
		}
	}

	result := multiply(matpow(T, n-2), base)

	var ans int64
	for i := 0; i < size; i++ {
		ans = (ans + result[i][0]) % mod
	}
	return int(ans)
}
