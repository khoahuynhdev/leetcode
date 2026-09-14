# 0836. Rectangle Overlap

[LeetCode Link](https://leetcode.com/problems/rectangle-overlap/)

Difficulty: Easy
Topics: Math, Geometry
Acceptance Rate: 48.3%

## Hints

### Hint 1

Resist the urge to enumerate the geometric configurations ("rec2 is above rec1", "rec2 is to the left", "rec2 is nested inside", ...). That path has a dozen cases and you will miss one. This is a Math/Geometry problem where the win comes from finding a single formula, not from a case analysis. Ask yourself: is it easier to characterize *overlap*, or its opposite?

### Hint 2

Axis-aligned rectangles are separable: the X dimension and the Y dimension are completely independent. A rectangle `[x1, y1, x2, y2]` is really the *product* of two intervals — `[x1, x2]` on the X axis and `[y1, y2]` on the Y axis. So the 2D question reduces to a 1D question asked twice: when do two intervals on a line share more than a single point?

### Hint 3

Two intervals `[a1, a2]` and `[b1, b2]` fail to overlap exactly when one ends at or before the other starts: `a2 <= b1 || b2 <= a1`. Flip that with De Morgan and you get the positive condition `a1 < b2 && b1 < a2`. The rectangles overlap iff *both* projections overlap, so:

```
rec1[0] < rec2[2] && rec2[0] < rec1[2] &&   // X intervals overlap
rec1[1] < rec2[3] && rec2[1] < rec1[3]      // Y intervals overlap
```

Note the strict `<`: that is precisely what makes touching edges and corners count as "no overlap".

## Approach

The key structural fact is **dimension separability**. An axis-aligned rectangle is the Cartesian product of an X interval and a Y interval, so the intersection of two rectangles is the Cartesian product of the two X-interval intersections and the two Y-interval intersections. The intersection area is positive if and only if *both* of those 1D intersections have positive length. One empty projection collapses the whole area to zero, no matter how generously the other dimension overlaps.

So the algorithm is:

1. Unpack `rec1 = [x1, y1, x2, y2]` into the X interval `[x1, x2]` and Y interval `[y1, y2]`; do the same for `rec2`.
2. Check the X projections for positive-length overlap: `rec1[0] < rec2[2] && rec2[0] < rec1[2]`.
3. Check the Y projections the same way: `rec1[1] < rec2[3] && rec2[1] < rec1[3]`.
4. Return the conjunction of the two.

Why is `a1 < b2 && b1 < a2` the right 1D test? It is easiest to derive by negation. Two intervals on a line miss each other in exactly two ways: `a` finishes before `b` starts (`a2 <= b1`), or `b` finishes before `a` starts (`b2 <= a1`). There is no third way — this is the classic separating-axis argument in one dimension. Negating that disjunction gives `a2 > b1 && b2 > a1`, which is the overlap test. The `<=` in the failure case becomes a strict `<` in the success case, which encodes "touching is not overlapping" for free.

An equivalent framing some people find more intuitive: the intersection interval is `[max(a1, b1), min(a2, b2)]`, and it is non-degenerate exactly when `max(a1, b1) < min(a2, b2)`. That expands to the same four comparisons.

Walking through Example 1 with `rec1 = [0,0,2,2]`, `rec2 = [1,1,3,3]`: X projections are `[0,2]` and `[1,3]`, and `0 < 3 && 1 < 2` holds, so they overlap on a segment of length 1. Y projections are identical by symmetry and also overlap. Both dimensions pass, so the answer is `true` — the shared region is the unit square `[1,2] x [1,2]`.

Example 2, `rec1 = [0,0,1,1]` and `rec2 = [1,0,2,1]`, is the instructive one. The Y projections are both `[0,1]` and overlap fully. But on X we have `[0,1]` and `[1,2]`: the test `rec2[0] < rec1[2]` is `1 < 1`, which is false. The rectangles share the vertical segment `x = 1` but that segment has zero width, so the area is zero and the answer is `false`. This is exactly the case that a naive `<=` would get wrong.

Example 3, `rec1 = [0,0,1,1]` and `rec2 = [2,2,3,3]`, fails on both dimensions — the rectangles are diagonally separated — and short-circuits to `false` on the very first comparison.

Honest note on difficulty: this problem is genuinely Easy once you see the separability insight, but the sub-50% acceptance rate is real and earned. Nearly all the failures come from two places — attempting case enumeration instead of finding the formula, and using `<=` where `<` is required. If you write the negation first and then flip it, both traps disappear.

## Complexity Analysis

Time Complexity: O(1) — a fixed four comparisons regardless of input, with short-circuit evaluation often doing fewer.
Space Complexity: O(1) — no allocation; the inputs are read in place.

## Edge Cases

- **Touching along an edge** (`[0,0,1,1]` and `[1,0,2,1]`): the projections meet at a single point in one dimension. This is the canonical off-by-one trap and the reason the comparisons must be strict `<` rather than `<=`. Zero-width intersection means zero area means no overlap.
- **Touching at exactly one corner** (`[0,0,1,1]` and `[1,1,2,2]`): *both* dimensions degenerate to a single shared point simultaneously. Still `false`, and the same strict inequality handles it without extra code.
- **Fully nested rectangles** (`[0,0,10,10]` containing `[2,2,3,3]`): a tempting hand-rolled case analysis based on corner containment can miss this, since neither rectangle's corners behave symmetrically. The interval formula handles containment with no special casing — `[2,3]` sits strictly inside `[0,10]` on both axes.
- **Identical rectangles**: overlap is total, so `true`. Worth checking that the strict inequalities do not accidentally reject it — they do not, because a valid rectangle has non-zero area and therefore `x1 < x2`.
- **Separated in one dimension only** (e.g. same X span, disjoint Y spans): a reminder that overlap requires the *conjunction*, not the disjunction, of the two projection tests. Getting this backwards is a common bug.
- **Negative and large coordinates** (down to `-10^9` and up to `10^9`): all four values can be negative, so do not assume a positive coordinate space. Go's `int` is 64-bit on all mainstream platforms, and since the solution only compares values and never adds or multiplies them, there is no overflow risk regardless.
