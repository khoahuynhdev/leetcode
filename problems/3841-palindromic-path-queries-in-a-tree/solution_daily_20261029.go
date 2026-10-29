package main

import (
	"strconv"
	"strings"
)

// Approach: a multiset of characters can form a palindrome iff at most one
// character has odd frequency, so encode each node's character as a bit and
// work with XOR parity masks.
//
// Let pre[x] be the XOR of the character bits on the root -> x path. Then the
// parity mask of the path u -> v is pre[u] ^ pre[v] ^ bit(lca(u, v)), because
// everything above the LCA cancels and the LCA itself cancels once too.
//
// Changing node x's character changes pre[y] for every y in x's subtree. With
// an Euler tour the subtree is a contiguous range [tin[x], tout[x]], so we use
// a Fenwick tree supporting range-XOR updates and point queries. LCA is found
// with binary lifting.
//
// Time: O((n + q) log n), Space: O(n log n).
func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}

	LOG := 1
	for (1 << LOG) < n {
		LOG++
	}
	up := make([][]int, LOG)
	for k := range up {
		up[k] = make([]int, n)
	}
	depth := make([]int, n)
	tin := make([]int, n)
	tout := make([]int, n)

	// Iterative DFS for Euler tour, parents and depths.
	timer := 0
	iter := make([]int, n)
	stack := []int{0}
	up[0][0] = 0
	tin[0] = timer
	timer++
	for len(stack) > 0 {
		u := stack[len(stack)-1]
		if iter[u] < len(adj[u]) {
			v := adj[u][iter[u]]
			iter[u]++
			// The root's parent is itself, so this skips both parent and root.
			if v == 0 || v == up[0][u] {
				continue
			}
			up[0][v] = u
			depth[v] = depth[u] + 1
			tin[v] = timer
			timer++
			stack = append(stack, v)
		} else {
			tout[u] = timer - 1
			stack = stack[:len(stack)-1]
		}
	}
	for k := 1; k < LOG; k++ {
		for v := 0; v < n; v++ {
			up[k][v] = up[k-1][up[k-1][v]]
		}
	}

	lca := func(a, b int) int {
		if depth[a] < depth[b] {
			a, b = b, a
		}
		diff := depth[a] - depth[b]
		for k := 0; diff > 0; k++ {
			if diff&1 == 1 {
				a = up[k][a]
			}
			diff >>= 1
		}
		if a == b {
			return a
		}
		for k := LOG - 1; k >= 0; k-- {
			if up[k][a] != up[k][b] {
				a = up[k][a]
				b = up[k][b]
			}
		}
		return up[0][a]
	}

	// Fenwick tree over XOR: range update via two point updates on the
	// difference array, point query via prefix XOR.
	bit := make([]int, n+2)
	pointXor := func(i, val int) {
		for i++; i <= n+1; i += i & -i {
			bit[i] ^= val
		}
	}
	rangeXor := func(l, r, val int) {
		pointXor(l, val)
		pointXor(r+1, val)
	}
	prefixMask := func(i int) int {
		res := 0
		for i++; i > 0; i -= i & -i {
			res ^= bit[i]
		}
		return res
	}

	chars := []byte(s)
	for v := 0; v < n; v++ {
		rangeXor(tin[v], tout[v], 1<<(chars[v]-'a'))
	}

	ans := []bool{}
	for _, q := range queries {
		parts := strings.Fields(q)
		u, _ := strconv.Atoi(parts[1])
		if parts[0] == "update" {
			c := parts[2][0]
			if c != chars[u] {
				rangeXor(tin[u], tout[u], 1<<(chars[u]-'a')^1<<(c-'a'))
				chars[u] = c
			}
			continue
		}
		v, _ := strconv.Atoi(parts[2])
		w := lca(u, v)
		mask := prefixMask(tin[u]) ^ prefixMask(tin[v]) ^ 1<<(chars[w]-'a')
		ans = append(ans, mask&(mask-1) == 0)
	}
	return ans
}
