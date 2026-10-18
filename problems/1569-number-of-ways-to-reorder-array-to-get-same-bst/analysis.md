# 1569. Number of Ways to Reorder Array to Get Same BST

[LeetCode Link](https://leetcode.com/problems/number-of-ways-to-reorder-array-to-get-same-bst/)

Difficulty: Hard
Topics: Array, Math, Divide and Conquer, Dynamic Programming, Tree, Union-Find, Binary Search Tree, Memoization, Combinatorics, Binary Tree, Fermat's Little Theorem
Acceptance Rate: 57.2%

## Hints

### Hint 1

Don't try to enumerate permutations. Think about which element *must* come first in any valid reordering, and what that forces about everything after it. This is a counting problem that breaks down recursively.

### Hint 2

The root is fixed as the first element. Every other value goes either to the left subtree (smaller than the root) or the right subtree (larger). What matters within each group is only its *relative order*. How do the two groups interact with each other in the array?

### Hint 3

Elements in the left subtree never affect where right-subtree elements land, and vice versa. So you may freely **interleave** the left sequence and the right sequence. The number of interleavings of sequences of lengths `L` and `R` is `C(L+R, L)`. Multiply that by the number of valid orderings of each subtree (recursively), and remember to subtract 1 for the original ordering.

## Approach

Let `ways(arr)` be the number of orderings of `arr` (including `arr` itself) that produce the same BST.

1. If `len(arr) <= 2`, there's exactly one ordering: the root must be first and the remaining element (if any) has nowhere else to go.
2. Otherwise, take `root = arr[0]`. Partition the rest into `left` (values `< root`) and `right` (values `> root`), keeping their original relative order. This relative order within each side is what determines the shape of each subtree.
3. Any merge of `left` and `right` that keeps each list's internal order produces the same tree, because insertions into the left subtree and the right subtree are independent. There are `C(len(left) + len(right), len(left))` such merges.
4. Each side can itself be reordered in `ways(left)` and `ways(right)` ways, so:

   `ways(arr) = C(n-1, len(left)) * ways(left) * ways(right)`

5. The answer is `ways(nums) - 1` (excluding the original array), taken modulo `1e9+7`.

For the binomial coefficients, precompute Pascal's triangle up to `n` with `C[i][j] = C[i-1][j-1] + C[i-1][j]` mod `1e9+7`. With `n ≤ 1000`, that's about 500k entries — cheap. (Alternatively, use factorials and modular inverses via Fermat's Little Theorem.)

**Example:** `nums = [3,4,5,1,2]`
- Root 3, `left = [1,2]`, `right = [4,5]`.
- `ways([1,2]) = 1`, `ways([4,5]) = 1`.
- `C(4, 2) = 6`, so `ways = 6`, answer `6 - 1 = 5`. ✅

## Complexity Analysis

Time Complexity: O(n^2) — Pascal's triangle is O(n^2), and the recursion does O(size) partitioning work per node, which is O(n^2) in the worst case (a skewed tree) and O(n log n) for balanced trees.
Space Complexity: O(n^2) — dominated by the Pascal table; the recursion adds O(n) stack depth and the partition slices along a root-to-leaf path are bounded by O(n^2) in the worst case.

## Edge Cases

- **Single element:** only one ordering exists, so the answer is 0 — make sure `ways - 1` isn't negative after the mod (use `(x - 1 + mod) % mod`).
- **Skewed tree (sorted ascending or descending):** every subtree has one side empty, so each split contributes `C(k, 0) = 1`; the answer is 0. This is also the deepest recursion (depth `n = 1000`), which is fine in Go.
- **Modular overflow:** multiply at most two values `< 1e9+7` at a time and reduce immediately; Go's 64-bit `int` holds the product safely.
- **Subtracting 1 when the product is 0 mod p:** the total count could be a multiple of `1e9+7`, so always add `mod` before taking the final modulo.
