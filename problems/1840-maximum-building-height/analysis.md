# 1840. Maximum Building Height

[LeetCode Link](https://leetcode.com/problems/maximum-building-height/)

Difficulty: Hard
Topics: Array, Math, Sorting
Acceptance Rate: 50.1%

## Hints

### Hint 1

`n` can be up to `10^9`, so you cannot build an array of heights and simulate every building. Notice that the number of *restrictions* is small (at most `10^5`). The interesting events only happen at the restricted buildings (plus building `1`, which is fixed at height `0`). Think about how to reason only about those "anchor" points and what happens *between* them.

### Hint 2

A restriction does not act in isolation. Because adjacent buildings differ by at most `1`, a low cap on one building forces nearby buildings to be low too. Sort the restrictions by index and ask: given building `a` is capped at height `Ha`, what is the tightest cap that propagates to a later building `b`? This is a relaxation/propagation idea — sweep left-to-right, then right-to-left, tightening each cap from its neighbors.

### Hint 3

Once every anchor has its *final, fully-propagated* maximum height, consider a single gap between two consecutive anchors `(idA, hA)` and `(idB, hB)`. Heights can climb up from both sides and meet at a peak in the middle. With `gap = idB - idA`, the peak you can reach is `(hA + hB + gap) / 2` (integer division — slopes of `+1` from each side meeting in the middle). The answer is the maximum peak over all consecutive gaps.

## Approach

We treat the problem as constraint propagation over a small set of "anchor" buildings.

**Step 1 — Collect anchors.**
The only buildings that matter are the restricted ones, building `1` (forced to height `0`), and building `n` (it can be the tallest, climbing freely on the right). Build a list of `(id, maxHeight)` points:

- Add `(1, 0)` — building 1 is forced to `0`.
- Keep every given restriction.
- If building `n` is not already restricted, add `(n, n-1)`. The cap `n-1` is the largest height physically reachable at index `n` (starting from `0` and rising by at most `1` per step), so it never over-constrains; it just gives us an anchor on the right edge.

Sort all anchors by `id`.

**Step 2 — Forward propagation (left to right).**
A building can be at most `1` taller than its left neighbor for each step between them. So for consecutive anchors `i-1` and `i`:

```
h[i] = min(h[i], h[i-1] + (id[i] - id[i-1]))
```

This pushes the unavoidable consequences of low caps (and the forced `0` at building 1) rightward.

**Step 3 — Backward propagation (right to left).**
Symmetrically, a cap on a later building limits earlier ones:

```
h[i] = min(h[i], h[i+1] + (id[i+1] - id[i]))
```

After both passes, every anchor holds its true maximum achievable height.

**Step 4 — Find the peak in each gap.**
Between two adjacent anchors `(idA, hA)` and `(idB, hB)` with `gap = idB - idA`, heights rise with slope `+1` from each side. The tallest point in the gap is:

```
peak = (hA + hB + gap) / 2   // integer division
```

The overall answer is the maximum of these peaks across all consecutive pairs. The anchor heights themselves are included automatically because a pair where one side is the taller anchor still yields at least that height.

**Worked example (Example 1):** `n = 5`, `restrictions = [[2,1],[4,1]]`.
Anchors after adding `(1,0)` and `(5,4)`: `[(1,0),(2,1),(4,1),(5,4)]`.
Forward pass leaves them `[0,1,1,2]`; backward pass keeps them `[0,1,1,2]`.
Gaps give peaks `1`, `(1+1+2)/2 = 2`, `(1+2+1)/2 = 2`. Maximum is **2**. ✅

This problem is genuinely a Hard — the propagation-both-ways plus the meet-in-the-middle peak formula is the crux. If you derived the `(hA + hB + gap) / 2` idea on your own, that is the hardest part done.

## Complexity Analysis

Let `m` be the number of restrictions.

Time Complexity: O(m log m) — dominated by sorting the anchors. The two propagation passes and the peak scan are each O(m).
Space Complexity: O(m) — for the list of anchor points.

## Edge Cases

- **No restrictions (`restrictions = []`):** Only anchors are `(1,0)` and the added `(n, n-1)`. The peak formula yields `n-1`, the unrestricted maximum (Example 2).
- **Building `n` already restricted:** Do not add a synthetic `(n, n-1)`; use the real cap, otherwise you would ignore a genuine constraint (Example 3 has `(10,3)`).
- **A restriction with `maxHeight` larger than reachable:** Forward/backward propagation tightens it down to what is physically achievable, so an overly generous cap never inflates the answer.
- **Out-of-order restrictions in the input:** Always sort by `id` first; the propagation logic assumes ascending indices.
- **Integer division in the peak formula:** `(hA + hB + gap) / 2` must use floor division — when the meeting point falls between two buildings, the achievable integer height is the floor.
- **Large `n` (up to `10^9`):** Never allocate per-building storage; everything is keyed off the `O(m)` anchors. Heights and indices fit comfortably in Go's `int` (64-bit on common platforms).
