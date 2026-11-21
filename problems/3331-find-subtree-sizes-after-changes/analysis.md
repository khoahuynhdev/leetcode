# 3331. Find Subtree Sizes After Changes

[LeetCode Link](https://leetcode.com/problems/find-subtree-sizes-after-changes/)

Difficulty: Medium
Topics: Array, Hash Table, String, Tree, Depth-First Search
Acceptance Rate: 62.3%

## Hints

### Hint 1

This is a two-phase problem: first figure out each node's *new* parent, then compute subtree sizes in the new tree. Think about which tree traversal naturally gives you "the current path from the root to this node."

### Hint 2

During a DFS, the recursion stack *is* the root-to-node path. You need to quickly answer "what is the closest node on this path with letter `c`?" There are only 26 letters — can you maintain something per letter as you enter and leave nodes?

### Hint 3

Keep 26 stacks (one per letter). On entering node `x`, the top of `stack[s[x]]` is its closest same-letter ancestor — that's the new parent. Push `x`, recurse into children, then pop on exit. Crucially, the new parent is always an ancestor in the *original* tree, so it comes before `x` in DFS preorder: walking the preorder backwards and adding each node's size into its new parent gives all final subtree sizes — no need to build or traverse the new tree.

## Approach

1. **Build children lists** from the `parent` array so we can DFS from the root.
2. **DFS with per-letter stacks.** Maintain `last[26]`, where `last[c]` holds the nodes with letter `c` on the current root-to-node path, in depth order.
   - On entering `x`: if `last[s[x]]` is non-empty, its top is the closest ancestor with the same letter, so `newParent[x] = top`. Otherwise `newParent[x] = parent[x]`.
   - Push `x` onto `last[s[x]]`, recurse into children, then pop.
   - Also record the preorder `order` of visited nodes.
   
   Because all reparenting is computed against the *original* tree (the changes happen simultaneously), this single DFS gives the correct new parent for every node.
3. **Accumulate sizes.** Every new parent is an ancestor of `x` in the original tree, so it appears earlier in the preorder. Iterate `order` from the end: set `answer[x] += 1` (itself) and then add `answer[x]` into `answer[newParent[x]]`. By the time a node is processed, all its new-tree descendants (which are all original-tree descendants appearing later in preorder) have already pushed their sizes into it.

**Example 2:** `parent = [-1,0,4,0,1]`, `s = "abbba"`. DFS path 0(a) → 1(b) → 4(a) → 2(b):
- Node 4 (`a`): `last[a] = [0]` → new parent 0.
- Node 2 (`b`): `last[b] = [1]` → new parent 1.
- Node 3 (`b`): path is 0 → 3, `last[b]` empty → stays under 0.

Final tree: 0 → {1, 3, 4}, 1 → {2}. Sizes `[5,2,1,1,1]`.

## Complexity Analysis

Time Complexity: O(n) — each node is visited once in the DFS and once in the accumulation pass; stack operations are O(1).
Space Complexity: O(n) — children lists, the per-letter stacks (total size ≤ depth), preorder, and recursion depth.

## Edge Cases

- **Single node (`n = 1`):** the answer is just `[1]`; make sure the root (parent `-1`) is never used as an index.
- **Chain of identical letters:** each node's closest same-letter ancestor is its own parent, so nothing changes.
- **All distinct letters:** no node is reparented; the answer equals original subtree sizes.
- **Simultaneous changes:** a node's new parent must be computed from the *original* tree, not a partially modified one (Example 2 shows why).
- **Deep trees (up to 10^5 nodes in a line):** recursion depth can be large; Go's growable goroutine stacks handle this, but in other languages you may need an iterative DFS.
