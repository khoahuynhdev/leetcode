# Segment Trees in Go

## What is a Segment Tree

A segment tree is a specialized binary tree data structure that efficiently handles range queries and updates on an array. Where a simple loop takes O(n) time to compute the sum or minimum of a range, and where updating an element in a prefix sum array takes O(n) time to rebuild, a segment tree does both operations in O(log n) time. This makes segment trees ideal for problems where you need to repeatedly query ranges and update values in any order.

The canonical segment tree problem is: given an array of integers, support two operations efficiently — update a single element at index `i` to a new value, and query the sum (or min, or max) of elements in range `[l, r]`. With a segment tree, both operations take O(log n) time, whereas naive approaches require O(n) for at least one operation.

Segment trees come up in competitive programming and technical interviews when you see phrases like "range sum queries with updates", "find the minimum in a subarray", or "count elements in a range satisfying some property". The key signal is that the problem requires both efficient querying and efficient updates — if it's just queries on a static array, prefix sums or sparse tables work fine. If it's just point updates with queries, a Binary Indexed Tree (Fenwick Tree) is simpler. Segment trees shine when you need the full flexibility of range queries with point or range updates.

## The Core Idea

A segment tree divides an array into segments stored as nodes in a binary tree. Each node represents a contiguous range of indices. The root represents the entire array `[0, n-1]`, and each node splits its range in half between its two children. Leaf nodes represent single elements. Each node stores an aggregate value for its range — the sum, minimum, maximum, or any other associative operation you care about.

```
Array: [1, 3, 5, 7, 9, 11]

Tree structure (each node shows [range: value]):
                    [0-5: 36]
                   /          \
            [0-2: 9]            [3-5: 27]
           /        \          /          \
       [0-1: 4]  [2: 5]    [3-4: 16]   [5: 11]
       /      \             /      \
   [0: 1]  [1: 3]      [3: 7]  [4: 9]
```

The height of the tree is O(log n), so any path from root to leaf takes O(log n) steps. Querying a range means combining the aggregate values from O(log n) nodes that together cover exactly the query range. Updating an element means walking from a leaf up to the root and recomputing aggregate values for O(log n) ancestors.

## Array-Based Representation

Segment trees are typically implemented using an array, not linked nodes. This uses a heap-like indexing scheme where if a node is at index `i`, its left child is at `2*i` and its right child is at `2*i + 1`. The root is at index 1, so we use 1-indexed arrays. This representation is cache-friendly and avoids pointer overhead.

For an array of size `n`, the segment tree array needs size `4*n`. This might seem wasteful, but it guarantees enough space for the full tree even in the worst case (when `n` is not a power of 2, some parts of the tree will have unused slots).

```go
type SegmentTree struct {
    tree []int // 1-indexed array: tree[i] stores aggregate for node i
    n    int   // size of the original array
}
```

## Building a Segment Tree

Building the tree is a post-order process. Recursively build the left and right children, then combine their results to get the parent's value. The base case is when the range is a single element — the node's value is just that element.

```go
func NewSegmentTree(arr []int) *SegmentTree {
    n := len(arr)
    st := &SegmentTree{
        tree: make([]int, 4*n),
        n:    n,
    }
    if n > 0 {
        st.build(arr, 1, 0, n-1)
    }
    return st
}

// build constructs the tree for node at index 'node' covering range [l, r]
func (st *SegmentTree) build(arr []int, node, l, r int) {
    if l == r {
        // Leaf node: single element
        st.tree[node] = arr[l]
        return
    }
    mid := l + (r-l)/2
    // Build left and right children
    st.build(arr, 2*node, l, mid)
    st.build(arr, 2*node+1, mid+1, r)
    // Combine results (for range sum)
    st.tree[node] = st.tree[2*node] + st.tree[2*node+1]
}
```

Building takes O(n) time because each element appears in exactly one leaf, and internal nodes are computed by combining children (no repeated work). The tree has O(n) nodes total.

## Point Update

To update a single element at index `i` to a new value, navigate to the leaf node representing index `i`, update it, then walk back up to the root updating each ancestor. Each ancestor's value is recomputed by combining its two children's values.

```go
// Update sets arr[idx] = val and propagates changes up the tree
func (st *SegmentTree) Update(idx, val int) {
    st.update(1, 0, st.n-1, idx, val)
}

func (st *SegmentTree) update(node, l, r, idx, val int) {
    if l == r {
        // Reached the leaf for idx
        st.tree[node] = val
        return
    }
    mid := l + (r-l)/2
    if idx <= mid {
        st.update(2*node, l, mid, idx, val)
    } else {
        st.update(2*node+1, mid+1, r, idx, val)
    }
    // Recompute this node's value from children
    st.tree[node] = st.tree[2*node] + st.tree[2*node+1]
}
```

Time complexity is O(log n) because you visit one node at each level of the tree.

## Range Query

To query the aggregate value (sum, min, max, etc.) over a range `[ql, qr]`, recursively break the range into segments that align with the tree's node ranges. Three cases arise at each node covering range `[l, r]`:

1. The query range completely contains the node's range `[l, r]` — return this node's stored value immediately.
2. The query range has no overlap with `[l, r]` — return the identity element (0 for sum, infinity for min, etc.).
3. Partial overlap — recurse into both children and combine their results.

```go
// Query returns the sum over range [ql, qr]
func (st *SegmentTree) Query(ql, qr int) int {
    return st.query(1, 0, st.n-1, ql, qr)
}

func (st *SegmentTree) query(node, l, r, ql, qr int) int {
    if qr < l || ql > r {
        // No overlap
        return 0 // identity for sum
    }
    if ql <= l && r <= qr {
        // Complete overlap: query range contains [l, r]
        return st.tree[node]
    }
    // Partial overlap: query both children
    mid := l + (r-l)/2
    left := st.query(2*node, l, mid, ql, qr)
    right := st.query(2*node+1, mid+1, r, ql, qr)
    return left + right
}
```

Time complexity is O(log n) because the query range can be decomposed into O(log n) disjoint segments that correspond to tree nodes. At each level of the tree, you touch at most four nodes (two boundary nodes and their siblings), so total nodes visited is proportional to tree height.

## Lazy Propagation

Standard segment tree point updates are O(log n), but what if you need to update an entire range efficiently? Naive approach would apply point updates to every element in the range — that's O(n log n) total, which defeats the purpose of using a segment tree. Lazy propagation is the technique that makes range updates O(log n).

The idea is to defer updates. When you need to update a range, instead of immediately updating all affected nodes, you mark nodes with a "lazy" tag indicating that this subtree needs an update. You only push the update down to children when you actually need to query or further update that subtree. This way, you do the minimum work necessary.

Each node gets a lazy value that represents a pending operation (e.g., "add 5 to all elements in this range"). When you access a node with a non-zero lazy value, you first "push down" the lazy value to its children before proceeding. This ensures correctness — lazy values are propagated just-in-time.

### Lazy Propagation Structure

```go
type LazySegmentTree struct {
    tree []int  // aggregate values (sum, min, max, etc.)
    lazy []int  // pending updates not yet pushed to children
    n    int
}

func NewLazySegmentTree(arr []int) *LazySegmentTree {
    n := len(arr)
    lst := &LazySegmentTree{
        tree: make([]int, 4*n),
        lazy: make([]int, 4*n),
        n:    n,
    }
    if n > 0 {
        lst.build(arr, 1, 0, n-1)
    }
    return lst
}

func (lst *LazySegmentTree) build(arr []int, node, l, r int) {
    if l == r {
        lst.tree[node] = arr[l]
        return
    }
    mid := l + (r-l)/2
    lst.build(arr, 2*node, l, mid)
    lst.build(arr, 2*node+1, mid+1, r)
    lst.tree[node] = lst.tree[2*node] + lst.tree[2*node+1]
}
```

### Push Down Mechanism

Before accessing a node's children, check if the node has a pending lazy value. If it does, apply the update to the current node, propagate the lazy value to the children, then clear the current node's lazy value.

```go
// pushDown applies pending updates from node to its children
func (lst *LazySegmentTree) pushDown(node, l, r int) {
    if lst.lazy[node] == 0 {
        return // no pending update
    }
    // Apply lazy update to this node's range [l, r]
    rangeSize := r - l + 1
    lst.tree[node] += lst.lazy[node] * rangeSize // for range-add on range sum

    if l != r {
        // Not a leaf: propagate lazy value to children
        lst.lazy[2*node] += lst.lazy[node]
        lst.lazy[2*node+1] += lst.lazy[node]
    }
    // Clear this node's lazy value
    lst.lazy[node] = 0
}
```

The exact formula for applying the lazy update to `tree[node]` depends on what aggregate you're tracking. For range sum with range add, you add `lazy[node] * rangeSize`. For range min with range add, you add `lazy[node]` to the min. For range assignment (set all elements to a value), you replace the aggregate entirely.

### Range Update

To add a value `val` to all elements in range `[ul, ur]`, recursively navigate the tree. At each node covering range `[l, r]`, push down any pending lazy updates, then:

1. If `[l, r]` is completely inside `[ul, ur]`, add `val` to this node's lazy value and apply it to the current node. Don't recurse further — you've deferred the work.
2. If `[l, r]` has no overlap with `[ul, ur]`, do nothing.
3. If partial overlap, recurse into both children, then recompute this node's aggregate from the updated children.

```go
// RangeAdd adds val to all elements in range [ul, ur]
func (lst *LazySegmentTree) RangeAdd(ul, ur, val int) {
    lst.rangeAdd(1, 0, lst.n-1, ul, ur, val)
}

func (lst *LazySegmentTree) rangeAdd(node, l, r, ul, ur, val int) {
    lst.pushDown(node, l, r)
    if ur < l || ul > r {
        // No overlap
        return
    }
    if ul <= l && r <= ur {
        // Complete overlap: apply lazy update
        lst.lazy[node] += val
        lst.pushDown(node, l, r) // apply immediately to this node
        return
    }
    // Partial overlap: recurse
    mid := l + (r-l)/2
    lst.rangeAdd(2*node, l, mid, ul, ur, val)
    lst.rangeAdd(2*node+1, mid+1, r, ul, ur, val)
    // Recompute this node from children (children are now up-to-date)
    lst.tree[node] = lst.tree[2*node] + lst.tree[2*node+1]
}
```

Time complexity is O(log n) because you mark O(log n) nodes with lazy values and push down O(log n) times.

### Range Query with Lazy Propagation

Querying with lazy propagation is nearly identical to the standard query, except you must push down lazy updates before using a node's value or recursing to children.

```go
// Query returns the sum over range [ql, qr]
func (lst *LazySegmentTree) Query(ql, qr int) int {
    return lst.query(1, 0, lst.n-1, ql, qr)
}

func (lst *LazySegmentTree) query(node, l, r, ql, qr int) int {
    if qr < l || ql > r {
        return 0 // no overlap
    }
    lst.pushDown(node, l, r) // ensure this node is up-to-date
    if ql <= l && r <= qr {
        // Complete overlap
        return lst.tree[node]
    }
    // Partial overlap
    mid := l + (r-l)/2
    left := lst.query(2*node, l, mid, ql, qr)
    right := lst.query(2*node+1, mid+1, r, ql, qr)
    return left + right
}
```

The key difference from the non-lazy version is the call to `pushDown` at the start. Without it, you might read stale values that don't reflect pending updates.

## Common Segment Tree Variations

Segment trees are a framework, not a single algorithm. The aggregate function and update operation determine what the tree can do. Here are the most common variations.

### Range Sum Queries with Point Updates

This is the simplest case: aggregate is sum, update changes a single element. Build function sums children, update walks to a leaf and recomputes ancestors, query combines sums from O(log n) segments. No lazy propagation needed.

```go
// Combine function for sum
st.tree[node] = st.tree[2*node] + st.tree[2*node+1]

// Identity for query
return 0
```

### Range Min/Max Queries with Point Updates

Instead of summing, each node stores the minimum (or maximum) of its range. Combine function is `min` or `max`, identity is infinity (or negative infinity for max).

```go
// Build and update: combine with min
st.tree[node] = min(st.tree[2*node], st.tree[2*node+1])

// Query: identity is infinity
if qr < l || ql > r {
    return math.MaxInt32
}
```

Used in problems like "range minimum query" where you need the smallest element in a range.

### Range Updates with Lazy Propagation

Add lazy propagation when you need to efficiently update entire ranges. Common operations are range-add (add a constant to all elements in a range) and range-assign (set all elements in a range to a constant). The lazy value represents the pending operation, and pushDown applies it.

For range-add on range-sum, the lazy value is the amount to add, and applying it means adding `lazy * rangeSize` to the sum.

For range-assign on range-sum, the lazy value is the new value for each element, and applying it means setting `tree[node] = lazy * rangeSize`. You also need to distinguish between "no pending update" and "set to 0", so use a sentinel or a separate boolean flag.

### Custom Queries: Finding Positions

Some problems require more than just aggregate values. For instance, you might need to find the leftmost index where a certain condition holds (e.g., first element greater than `x`, or first element equal to zero). You can augment the segment tree to answer these queries by making the query function smarter — instead of just combining values, it can traverse the tree more selectively.

Problem 3721 (Longest Balanced Subarray II) uses this technique. The segment tree tracks the minimum and maximum value in each range, and a custom `findLeftmostZero` function walks the tree looking for the leftmost range containing zero, then recurses into that subtree until it finds the exact index.

```go
// Custom query: find leftmost index where tree[node] contains 0
func (lst *LazySegmentTree) findLeftmostZero(node, l, r int) int {
    lst.pushDown(node, l, r)
    if lst.tree[node].mn > 0 || lst.tree[node].mx < 0 {
        // Range [l, r] does not contain 0
        return -1
    }
    if l == r {
        // Leaf node and range contains 0, so this is the answer
        return l
    }
    mid := l + (r-l)/2
    // Try left child first (leftmost)
    if result := lst.findLeftmostZero(2*node, l, mid); result != -1 {
        return result
    }
    // Otherwise try right child
    return lst.findLeftmostZero(2*node+1, mid+1, r)
}
```

This is a powerful pattern. By storing not just a single aggregate but a richer summary (min and max, or a set of values, or counts of different types), you can answer complex queries in O(log n) time.

## When to Use Segment Trees vs Alternatives

Segment trees are not the only data structure for range queries. Understanding when to use which tool is important for both solving problems efficiently and for interviews where you're expected to justify your choice.

**Prefix Sums**: If the array is static (no updates), prefix sums solve range sum queries in O(1) time with O(n) preprocessing. They're simpler and faster than segment trees. Use prefix sums when you only need to query a static array. If updates are required, prefix sums break down (each update takes O(n) to rebuild).

**Binary Indexed Tree (Fenwick Tree)**: BIT is a simpler alternative to segment trees for certain problems. It supports point update and prefix sum query (sum from index 0 to `i`) in O(log n) time. From prefix sums you can compute any range sum in O(1) as `sum[r] - sum[l-1]`. BIT uses less memory than segment trees and has a simpler implementation, but it only works for operations where the inverse exists (sum, XOR) — you can't use it for min/max without more complex tricks. BIT also doesn't naturally support range updates; you need difference arrays or other workarounds for that. If your problem is "point update, range sum query", prefer BIT. If you need range min/max, or range updates, use a segment tree.

**Sparse Table**: For range min/max queries on a static array, sparse tables give O(1) query time after O(n log n) preprocessing. They're ideal when you need millions of queries on a static array. They don't support updates at all.

**Segment Trees**: Use segment trees when you need flexible range queries (sum, min, max, GCD, or custom aggregates) with updates (point or range). Segment trees are the most general structure and handle all of these cases efficiently. The tradeoff is more code complexity and slightly worse constant factors. If you need lazy propagation for range updates, segment trees are essentially the only option.

In interviews, segment trees signal that you know advanced data structures, but they're overkill for simpler problems. If prefix sums or BIT suffice, mention them first to show you understand tradeoffs. If the problem requires range updates or non-invertible aggregates, segment trees are the right tool.

## Tips and Tricks

Segment trees have a reputation for being tricky to implement correctly. Here are practical tips to avoid common pitfalls and make your code cleaner.

**Array size should be 4n**. You might see some sources say 2n is enough if n is a power of 2, but using 4n is the safe, simple rule that always works. It wastes a bit of space but avoids hard-to-debug indexing errors. For interview problems with constraints like n ≤ 100,000, the memory difference is negligible.

**Use 1-indexed arrays for tree nodes**. This makes the parent-child relationship clean: left child is `2*i`, right child is `2*i+1`, parent is `i/2`. If you use 0-indexed, the formulas become `2*i+1` and `2*i+2`, which is more error-prone. Initialize `tree := make([]int, 4*n)` and never use index 0 for a node — it's wasted but simplifies everything else.

**Write a combine function**. Instead of scattering `tree[node] = tree[2*node] + tree[2*node+1]` throughout your code, write a helper function that combines two child values. This makes it easier to change the aggregate type (sum to min to max) and reduces copy-paste errors.

```go
func (st *SegmentTree) combine(left, right int) int {
    return left + right // or min(left, right), or max(left, right)
}

st.tree[node] = st.combine(st.tree[2*node], st.tree[2*node+1])
```

**Push down before everything**. If you're using lazy propagation, call `pushDown` at the very start of every function that accesses a node (update, query, custom operations). It's easy to forget, and the bug manifests as wrong answers on test cases that trigger lazy values. Make it a habit: first line of any node access is `pushDown`.

**Test with small examples first**. Segment trees are hard to debug when they go wrong. Before submitting, test on a tiny array (size 3 or 4) where you can manually trace every node and verify the tree structure matches your expectations. Print the tree array after building it and check that it makes sense.

**Recognize segment tree problems**. In interviews, segment trees are rarely stated explicitly. Look for these clues: problem requires both queries and updates, multiple queries on ranges, updates are range-based or frequent, naive solution times out. Phrases like "for each query, compute the sum/min/max of elements from L to R" combined with "you may update element at index X" are strong signals. If you see this, immediately consider segment trees or BIT.

**Lazy propagation is for range updates only**. Don't use lazy propagation if all updates are point updates — it adds complexity with no benefit. Lazy propagation shines when you have operations like "add 5 to all elements from index L to R", which occur frequently. If point updates are sufficient, stick to the simpler non-lazy version.

**Be careful with lazy value semantics**. The meaning of `lazy[node]` depends on the operation. For range-add, it's the amount to add. For range-assign, it's the value to set. If you need both operations, you might need two lazy arrays or a more complex lazy structure. Mixing operations requires careful thought about how to combine pending updates when pushing down.

**Watch out for integer overflow**. When doing range sums, the sum can exceed the range of `int`. Use `int64` for the tree array if the problem's constraints allow large sums. This is a common source of wrong answers on problems that look like they work but fail on large inputs.

**Segment trees for 2D ranges**. You can build a segment tree where each node contains another segment tree, creating a 2D structure. This is rare and complex, but it allows range queries on rectangles in O(log² n) time. In most interviews, 2D range queries are solved with 2D prefix sums or other techniques, but if you need updates, nested segment trees are an option.

## Complexity Reference

| Operation | Time Complexity | Space Complexity |
|-----------|----------------|------------------|
| Build tree | O(n) | O(n) |
| Point update | O(log n) | - |
| Range query | O(log n) | - |
| Range update (with lazy propagation) | O(log n) | - |
| Point query | O(log n) | - |

The space complexity is O(n) for the tree array (technically 4n elements, but asymptotically O(n)). If using lazy propagation, you need another O(n) array for lazy values. All operations are O(log n) because the tree height is log n.

Compared to naive approaches: a simple loop for range query is O(n), updating and recomputing prefix sums is O(n), so segment trees provide a logarithmic speedup for both. This matters when you have many queries and updates — for Q queries and U updates, naive is O(Q * n + U * n), while segment trees are O((Q + U) * log n).

## Problems Log

| Problem | Difficulty | Key Insight | Solution |
|---------|-----------|-------------|----------|
| [307. Range Sum Query - Mutable](https://leetcode.com/problems/range-sum-query-mutable/) | Medium | Classic segment tree application: point update + range sum query. Build tree with sum aggregates, update propagates changes up ancestors, query combines disjoint segments. Can also use Binary Indexed Tree. | - |
| [3721. Longest Balanced Subarray II](https://leetcode.com/problems/longest-balanced-subarray-ii/) | Hard | Use sweep line with balance tracking. For each right endpoint r, maintain balance[l] = distinct_even(l,r) - distinct_odd(l,r) for all left endpoints. When processing nums[r], if previous occurrence is at prev, range-add +1 (even) or -1 (odd) to [prev+1, r]. Answer is longest subarray where balance[l] = 0. Segment tree with lazy propagation for range updates + custom findLeftmostZero query (traverse tree looking for range containing 0, recurse into left child first for leftmost). Store min and max in each node to quickly determine if range contains zero. | [solution.go](../problems/3721-longest-balanced-subarray-ii/solution.go) |
| [218. The Skyline Problem](https://leetcode.com/problems/the-skyline-problem/) | Hard | Can use segment tree with range max queries and range updates to track building heights. Alternative approaches use sweep line with priority queue. Segment tree is overkill here but illustrates its flexibility. | - |
| [308. Range Sum Query 2D - Mutable](https://leetcode.com/problems/range-sum-query-2d-mutable/) | Hard | 2D segment tree (each node contains a 1D segment tree) or Binary Indexed Tree in 2D. For interviews, BIT is usually simpler unless range updates are needed. | - |
| [327. Count of Range Sum](https://leetcode.com/problems/count-of-range-sum/) | Hard | Count subarrays with sum in [lower, upper]. Use prefix sums, then for each prefix sum, query how many previous prefix sums fall in the required range. Segment tree on coordinate-compressed prefix sums with point updates and range count queries. Merge sort approach is simpler but segment tree is more flexible. | - |
| [699. Falling Squares](https://leetcode.com/problems/falling-squares/) | Hard | Track max height in ranges with range updates. Use segment tree with lazy propagation for range-set (assign all elements in range to a value) and range-max query. Coordinate compression needed since coordinates can be large. | - |

The most important problem to understand is 307 (Range Sum Query - Mutable), which is the canonical segment tree problem and tests your ability to implement the core structure correctly. Problem 3721 is an excellent example of augmenting segment trees with custom queries — it's not just about sums or mins, but about finding specific properties in ranges, which requires thinking beyond the standard template.
