# 2196. Create Binary Tree From Descriptions

[LeetCode Link](https://leetcode.com/problems/create-binary-tree-from-descriptions/)

Difficulty: Medium
Topics: Array, Hash Table, Tree, Binary Tree
Acceptance Rate: 82.3%

## Hints

### Hint 1

The input gives you edges (parent → child relationships) scattered in arbitrary
order, but you need to assemble an actual linked tree structure. Think about how
you can look up or create a node by its *value* quickly, no matter when you first
encounter it. Which data structure lets you find "the node for value `v`" in
constant time?

### Hint 2

A hash map from value → `*TreeNode` is your friend. As you scan the descriptions,
you will repeatedly need the node objects for both `parent` and `child`. Create
them on first sight and reuse them afterward, so that wiring up `parent.Left` or
`parent.Right` connects to the *same* shared node every time it is referenced.

### Hint 3

The only remaining puzzle is: which node is the root? The root is the single node
that is **never a child** of anyone. So while you process the edges, keep a set of
all values that appear in the `child` position. After the pass, the root is the
one value that exists in your node map but is *not* in the child set. With a hash
map plus a child-set, the whole thing is one linear pass to build, and one scan to
find the root.

## Approach

The problem hands you a flat list of parent/child edges and asks you to rebuild
the binary tree they describe. The challenge is purely bookkeeping: the same value
can appear as a parent in one description and as a child in another, in any order,
so we need a way to refer to "the node for value `v`" consistently.

**Step 1 — Map values to nodes.**
Maintain `nodes`, a `map[int]*TreeNode`. Define a small helper that returns the
node for a value, creating and inserting it on first use. This guarantees there is
exactly one `*TreeNode` per value, and any `Left`/`Right` pointer we set refers to
that shared instance.

**Step 2 — Wire up the edges.**
For each `[parent, child, isLeft]`:
- Fetch (or create) the parent node and the child node via the helper.
- If `isLeft == 1`, set `parent.Left = child`, otherwise `parent.Right = child`.
- Record `child` in a `children` set (`map[int]bool`) so we know it has a parent.

**Step 3 — Find the root.**
A valid binary tree has exactly one root, and the root is the only value that is
never listed as someone's child. Scan the keys of `nodes`; the value that is *not*
present in `children` is the root. Return its node.

*Walkthrough on Example 1:*
`[[20,15,1],[20,17,0],[50,20,1],[50,80,0],[80,19,1]]`. Processing the edges
creates nodes for 20, 15, 17, 50, 80, 19 and links them: 20→(15 left, 17 right),
50→(20 left, 80 right), 80→(19 left). The `children` set is {15, 17, 20, 80, 19}.
The only value in `nodes` but not in `children` is **50**, so 50 is the root —
matching the expected output.

This is the optimal approach: every description is touched a constant number of
times, and a hash map gives O(1) lookups.

## Complexity Analysis

Time Complexity: O(n), where n is the number of descriptions. We do one pass to
build the tree and one pass over the node map (also O(n) distinct values) to find
the root.

Space Complexity: O(n) for the node map, the children set, and the tree nodes
themselves.

## Edge Cases

- **Single node tree (no descriptions never happens, but a single edge):** With
  the smallest valid input (one description), two nodes are created and one is
  clearly the root — the helper still creates both correctly. The constraints
  guarantee `descriptions.length >= 1`, so the map is never empty.
- **A node appears as a child before it appears as a parent (or vice versa):**
  Because the helper creates nodes lazily and reuses them, the order in which
  values are encountered does not matter. This is the key correctness concern.
- **Reusing the same node object:** If you accidentally create a fresh node each
  time instead of reusing from the map, subtrees get dropped. The map prevents
  this.
- **Distinguishing left vs right child:** Mixing up `isLeft == 1` (left) and
  `isLeft == 0` (right) silently produces a wrong-but-valid-looking tree, so the
  conditional must be exact.
