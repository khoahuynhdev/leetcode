# 3841. Palindromic Path Queries in a Tree

[LeetCode Link](https://leetcode.com/problems/palindromic-path-queries-in-a-tree/)

Difficulty: Hard
Topics: Array, String, Divide and Conquer, Bit Manipulation, Tree, Depth-First Search, Binary Indexed Tree, Segment Tree
Acceptance Rate: 59.9%

## Hints

### Hint 1

Forget the tree for a moment. When can a bag of letters be rearranged into a palindrome? You don't need the exact counts. You only need to know whether each count is **odd or even**. With 26 letters, that information fits in a single integer.

### Hint 2

Represent each node's letter as one bit (`1 << (c - 'a')`) and combine letters with XOR. Now think of **prefix sums on a tree**: define `pre[x]` as the XOR of all bits on the path from the root to `x`. How can you get the mask of an arbitrary path `u -> v` from `pre[u]`, `pre[v]`, and the **lowest common ancestor**?

### Hint 3

`mask(u, v) = pre[u] ^ pre[v] ^ bit(lca(u, v))`. Everything above the LCA cancels out, and the LCA cancels once too, so you add it back.

Updates are the hard part. Changing node `x`'s letter changes `pre[y]` for **every node in x's subtree**. With an Euler tour (DFS entry/exit times), a subtree is a **contiguous range**. So you need "XOR a value onto a range" and "read a single position". A Fenwick tree over a difference array handles both in `O(log n)`.

## Approach

1. **Palindrome test with a bitmask.** A multiset of letters can form a palindrome iff at most one letter has an odd count. If `mask` holds the parity of each letter, the test is `mask & (mask - 1) == 0`, meaning zero or one bit is set.

2. **Root the tree at node 0 and do a DFS.** Record:
   - `tin[x]` / `tout[x]`: Euler-tour entry time and the last entry time inside x's subtree. Subtree of `x` = positions `[tin[x], tout[x]]`.
   - `depth[x]` and the binary-lifting table `up[k][x]` for LCA queries.

   The DFS is iterative, so a long chain (n = 5·10⁴) can't overflow the stack.

3. **Prefix masks in a Fenwick tree.** Conceptually `pre[y]` = XOR of `bit(s[a])` over every ancestor `a` of `y`, including `y` itself. Node `a` contributes to exactly the nodes in its subtree, so for every node `a` we XOR `bit(s[a])` onto the range `[tin[a], tout[a]]`.
   Range XOR + point query is the classic Fenwick difference trick: XOR `val` at `l` and at `r + 1`. The prefix XOR up to `i` then gives the value at position `i`.

4. **Update `x` to `c`:** the delta is `bit(old) ^ bit(c)`. XOR it onto `[tin[x], tout[x]]`. If the letter doesn't change, skip it.

5. **Query `u v`:** compute `w = lca(u, v)`, then
   `mask = point(tin[u]) ^ point(tin[v]) ^ bit(s[w])`, and answer `mask & (mask - 1) == 0`.

**Example (Example 1):** chain `0 - 1 - 2`, `s = "aac"`.
`pre[0] = a`, `pre[1] = a^a = 0`, `pre[2] = c`. For query `0 2`, the LCA is 0:
`mask = a ^ c ^ a = c`, so one odd letter means true.
After `update 1 b`, nodes 1 and 2 (the subtree of 1) get XORed with `a ^ b`, so `pre[2] = a^b^c`.
`mask = a ^ (a^b^c) ^ a = a^b^c`, so three odd letters means false.

**Why it works:** XOR is its own inverse, so shared ancestors cancel exactly like subtracting prefix sums. The Euler tour turns "all descendants" into "a contiguous interval", which is what makes a 1D Fenwick tree enough.

## Complexity Analysis

Time Complexity: O((n + q) log n). Building the lifting table and initializing the Fenwick tree take O(n log n). Each update or query takes O(log n).
Space Complexity: O(n log n) for the binary-lifting table. Everything else is O(n).

## Edge Cases

- **Single node (`n = 1`, no edges):** any query is `0 0`, a single letter, which is always a palindrome. The lifting table and DFS must still work with no edges.
- **`u == v`:** the path is one node. The formula gives `pre[u] ^ pre[u] ^ bit(s[u]) = bit(s[u])`, which correctly has one bit set.
- **One endpoint is the ancestor of the other:** then the LCA is an endpoint. The same formula holds, but it's a common place for off-by-one mistakes (counting the LCA twice or zero times).
- **Updating a node to its current letter:** the delta is 0. Skipping it is an optimization, but make sure the stored letter isn't corrupted.
- **Updating the root:** this affects the whole Euler range `[0, n - 1]`, so the Fenwick tree must be sized to allow writing at `r + 1 = n`.
- **Deep chain trees:** recursion depth can reach 5·10⁴. Use an iterative DFS (or be confident about your language's stack limits).
- **Only updates, no queries:** return an empty array, not `nil`/`null`.
