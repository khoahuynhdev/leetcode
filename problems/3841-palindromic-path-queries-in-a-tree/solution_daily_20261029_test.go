package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		s        string
		queries  []string
		expected []bool
	}{
		{
			"example 1: line tree with update breaking palindrome",
			3, [][]int{{0, 1}, {1, 2}}, "aac",
			[]string{"query 0 2", "update 1 b", "query 0 2"},
			[]bool{true, false},
		},
		{
			"example 2: star tree with multiple updates",
			4, [][]int{{0, 1}, {0, 2}, {0, 3}}, "abca",
			[]string{"query 1 2", "update 0 b", "query 2 3", "update 3 a", "query 1 3"},
			[]bool{false, false, true},
		},
		{
			"edge case: single node tree",
			1, [][]int{}, "z",
			[]string{"query 0 0", "update 0 a", "query 0 0"},
			[]bool{true, true},
		},
		{
			"edge case: update to the same character is a no-op",
			2, [][]int{{0, 1}}, "ab",
			[]string{"query 0 1", "update 1 a", "query 0 1", "update 1 a", "query 1 0"},
			[]bool{false, true, true},
		},
		{
			"edge case: chain with sub-path queries",
			5, [][]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}}, "abcba",
			[]string{"query 0 4", "query 1 3", "query 0 3", "update 2 a", "query 0 3"},
			[]bool{true, true, false, true},
		},
		{
			"edge case: LCA strictly inside the path and root update",
			5, [][]int{{0, 1}, {1, 2}, {1, 3}, {3, 4}}, "aabba",
			[]string{"query 2 4", "query 0 4", "update 0 b", "query 0 4", "query 2 2"},
			[]bool{true, false, true, true},
		},
		{
			"edge case: only updates, no queries",
			2, [][]int{{0, 1}}, "ab",
			[]string{"update 0 c", "update 1 d"},
			[]bool{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := palindromePath(tt.n, tt.edges, tt.s, tt.queries)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}

// bruteForce walks each path explicitly and counts character parities.
func bruteForce(n int, edges [][]int, s string, queries []string) []bool {
	adj := make([][]int, n)
	for _, e := range edges {
		adj[e[0]] = append(adj[e[0]], e[1])
		adj[e[1]] = append(adj[e[1]], e[0])
	}
	chars := []byte(s)
	ans := []bool{}
	for _, q := range queries {
		var op string
		var u int
		var rest string
		fmt.Sscanf(q, "%s %d %s", &op, &u, &rest)
		if op == "update" {
			chars[u] = rest[0]
			continue
		}
		var v int
		fmt.Sscanf(rest, "%d", &v)
		parent := make([]int, n)
		for i := range parent {
			parent[i] = -2
		}
		parent[u] = -1
		queue := []int{u}
		for len(queue) > 0 {
			x := queue[0]
			queue = queue[1:]
			for _, y := range adj[x] {
				if parent[y] == -2 {
					parent[y] = x
					queue = append(queue, y)
				}
			}
		}
		mask := 0
		for x := v; x != -1; x = parent[x] {
			mask ^= 1 << (chars[x] - 'a')
		}
		ans = append(ans, mask&(mask-1) == 0)
	}
	return ans
}

func TestSolutionRandomized(t *testing.T) {
	rng := rand.New(rand.NewSource(3841))
	for iter := 0; iter < 300; iter++ {
		n := rng.Intn(12) + 1
		edges := make([][]int, 0, n-1)
		for v := 1; v < n; v++ {
			edges = append(edges, []int{rng.Intn(v), v})
		}
		sb := make([]byte, n)
		for i := range sb {
			sb[i] = byte('a' + rng.Intn(3))
		}
		queries := make([]string, rng.Intn(20)+1)
		for i := range queries {
			if rng.Intn(2) == 0 {
				queries[i] = fmt.Sprintf("update %d %c", rng.Intn(n), 'a'+rng.Intn(3))
			} else {
				queries[i] = fmt.Sprintf("query %d %d", rng.Intn(n), rng.Intn(n))
			}
		}
		got := palindromePath(n, edges, string(sb), queries)
		want := bruteForce(n, edges, string(sb), queries)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("n=%d edges=%v s=%s queries=%v: got %v, want %v", n, edges, sb, queries, got, want)
		}
	}
}
