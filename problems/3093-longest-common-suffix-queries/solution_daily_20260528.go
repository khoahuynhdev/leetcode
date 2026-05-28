// 3093. Longest Common Suffix Queries
//
// Build a trie over the reversed wordsContainer entries. At every node store
// the (shortest, then earliest) container index that passes through it. For
// each query, walk the trie following the reversed query characters until a
// child is missing; the current node's stored index is the answer.
//
// Inserting words in input order with a strict "<" length comparison makes the
// shortest-then-earliest tiebreak fall out automatically. The root node holds
// the global best, which covers queries that share no suffix at all.

package main

type suffixTrieNode struct {
	children [26]*suffixTrieNode
	bestIdx  int
	bestLen  int
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	root := &suffixTrieNode{bestIdx: -1}

	for i, w := range wordsContainer {
		node := root
		if node.bestIdx == -1 || len(w) < node.bestLen {
			node.bestIdx = i
			node.bestLen = len(w)
		}
		for j := len(w) - 1; j >= 0; j-- {
			c := w[j] - 'a'
			child := node.children[c]
			if child == nil {
				child = &suffixTrieNode{bestIdx: i, bestLen: len(w)}
				node.children[c] = child
			} else if len(w) < child.bestLen {
				child.bestIdx = i
				child.bestLen = len(w)
			}
			node = child
		}
	}

	ans := make([]int, len(wordsQuery))
	for i, q := range wordsQuery {
		node := root
		for j := len(q) - 1; j >= 0; j-- {
			c := q[j] - 'a'
			if node.children[c] == nil {
				break
			}
			node = node.children[c]
		}
		ans[i] = node.bestIdx
	}

	return ans
}
