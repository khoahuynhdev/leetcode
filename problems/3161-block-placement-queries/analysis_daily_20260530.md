# 3161. Block Placement Queries

[LeetCode Link](https://leetcode.com/problems/block-placement-queries/)

Difficulty: Hard
Topics: Array, Binary Search, Binary Indexed Tree, Segment Tree
Acceptance Rate: 24.4%

## Hints

### Hint 1

For any query `[2, x, sz]`, the only thing that matters about the obstacles is the set of *gaps* between consecutive obstacles inside `[0, x]`, plus the leftover space between the largest obstacle ≤ `x` and `x` itself. A block of size `sz` fits if and only if one of those gaps is at least `sz`. So the real question is: how do you maintain "the maximum gap between consecutive obstacles" while obstacles are being inserted one at a time?

### Hint 2

Inserting an obstacle changes at most two gaps: the gap landing at the new obstacle, and the gap landing at its successor (the previous "long" gap gets split into two). That locality is what makes a segment tree feasible. Coordinate-compress every position you will ever see (all `x` values from both query types, plus a sentinel at 0), then index a segment tree by compressed position. Pair it with a way to find the predecessor and successor of any position among the currently-placed obstacles.

### Hint 3

Use **two** segment trees over the compressed positions:

1. A *presence* tree (sum / "exists" semantics) that lets you descend to find the rightmost obstacle ≤ index `i` (predecessor) and leftmost obstacle ≥ index `i` (successor).
2. A *gap* tree (range-max) where `gap[i]` stores the distance from the obstacle at compressed index `i` to its predecessor obstacle. Inserting `x` updates `gap[x] = x - pred` and `gap[succ] = succ - x`.

Treat position `0` as a permanent sentinel obstacle so the very first real gap (from origin to the first obstacle) is handled by the same formula. A query `[2, x, sz]` then computes `max(rangeMax(gap, 1..idx(pred)), x - pred) ≥ sz`, where `pred` is the largest obstacle position ≤ `x`.

## Approach

**Setup.** Collect every distinct position you will ever touch — all `x` values from queries of both types and the sentinel `0` — then sort and assign each a compressed index. Build two segment trees over those indices:

- `presence[i]` ∈ {0, 1} tells whether an obstacle is currently at the position with compressed index `i`. Internal nodes hold the sum, but more importantly we use the tree to *descend* to the rightmost / leftmost `1` in a range.
- `gap[i]` stores, for each currently active obstacle, the distance to its immediate predecessor obstacle (with `0` acting as the predecessor of the leftmost real obstacle). Internal nodes hold the max.

Mark the sentinel: `presence[idx(0)] = 1`, `gap[idx(0)] = 0`.

**Type 1 — insert obstacle at `x`.** Let `xi = idx(x)`.

1. `predIdx = rightmost 1 in presence[0..xi-1]` — guaranteed to exist thanks to the sentinel. Let `predPos = positions[predIdx]`.
2. `succIdx = leftmost 1 in presence[xi+1..n-1]`, possibly `-1` if none.
3. Set `presence[xi] = 1`.
4. Set `gap[xi] = x - predPos` (the new gap ending at `x`).
5. If `succIdx ≠ -1`, set `gap[succIdx] = positions[succIdx] - x` (the old gap, which used to span from `predPos` to `positions[succIdx]`, is now shortened on the left).

Each insert touches O(log n) nodes in each tree.

**Type 2 — query `(x, sz)`.** Let `xi = idx(x)`.

1. `predIdx = rightmost 1 in presence[0..xi]` (note: includes `xi` itself in case `x` *is* an obstacle).
2. `predPos = positions[predIdx]`.
3. `interiorMax = rangeMax(gap, 1..predIdx)` — we skip index 0 (the sentinel's gap is 0 by convention).
4. `trailing = x - predPos` — the leftover space from the last obstacle ≤ `x` up to `x`.
5. Answer is `max(interiorMax, trailing) ≥ sz`.

The `trailing` term matters because the segment tree only stores gaps that *end at an obstacle*; the gap from the last obstacle to `x` has no obstacle on its right and isn't stored anywhere, so we add it explicitly.

**Worked example (Example 2).** `queries = [[1,7],[2,7,6],[1,2],[2,7,5],[2,7,6]]`. Compressed positions are `[0, 2, 7]`.

- `[1,7]` → predecessor of 7 is 0; insert. `gap[idx(7)] = 7`.
- `[2,7,6]` → pred = 7, interior max over `gap` indices `[1, idx(7)]` = 7, trailing = 0. `7 ≥ 6` → **true**.
- `[1,2]` → predecessor of 2 is 0, successor is 7. Insert. `gap[idx(2)] = 2`, `gap[idx(7)] = 7 - 2 = 5`.
- `[2,7,5]` → pred = 7, interior max = `max(2, 5) = 5`, trailing = 0. `5 ≥ 5` → **true**.
- `[2,7,6]` → same, `5 ≥ 6` → **false**.

Result: `[true, true, false]`. ✓

This is genuinely a hard problem — the trick of *coordinate-compressing query positions too* (not just obstacle positions) is what makes the predecessor query fast, and the locality observation about gaps is what keeps inserts cheap. If you got stuck thinking the obstacles needed to be maintained in a balanced BST or skip list, that intuition is right; coordinate compression + a segment tree just gives you the same power without writing a BST in Go.

## Complexity Analysis

Time Complexity: O((Q + P) log P), where `Q` is the number of queries and `P` is the number of distinct positions across all queries (`P ≤ Q + 1`). Each query performs a constant number of segment-tree operations, each O(log P).

Space Complexity: O(P) for the compressed-position table and the two segment trees.

## Edge Cases

- **No obstacle ≤ `x`.** The sentinel at 0 still acts as the predecessor, so `predPos = 0`, `trailing = x`, and the answer reduces to `x ≥ sz` — exactly the empty-line case.
- **`x` itself is an obstacle.** Predecessor lookup includes `xi`, so `predPos = x` and `trailing = 0`. The block has to live entirely in one of the interior gaps; we don't double-count anything past `x`.
- **First obstacle inserted.** No successor exists, so step 5 of insert is skipped. The first real gap (origin to obstacle) is captured cleanly because `predPos` is the sentinel `0`.
- **Inserting between two existing obstacles.** Both `gap[xi]` and `gap[succIdx]` are rewritten in the same call; the old combined gap effectively disappears from the max tree (replaced by the smaller of its two halves).
- **Query for a tiny `sz`.** Even `sz = 1` always passes whenever `x ≥ 1` and at least one gap (including trailing) is ≥ 1 — typically trivially true.
- **Query `x` equals a position never used by a type-1 query.** Coordinate compression must include all `x` values from type-2 queries too, otherwise `idx(x)` would be undefined. This is why we collect positions from both query types.
