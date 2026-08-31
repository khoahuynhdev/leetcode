# 2058. Find the Minimum and Maximum Number of Nodes Between Critical Points

[LeetCode Link](https://leetcode.com/problems/find-the-minimum-and-maximum-number-of-nodes-between-critical-points/)

Difficulty: Medium
Topics: Linked List
Acceptance Rate: 70.5%

## Hints

### Hint 1

Being a critical point is a purely *local* property: a node qualifies based only on itself and its two immediate neighbours. That means you never need to look at the whole list at once, and you never need random access. Ask yourself what the smallest window of the list you must hold at any moment is, and whether one forward pass can carry that window along.

### Hint 2

Walk the list with three references — previous, current, next — and a running index counter, so each node knows its own position. Whenever `current` is strictly greater than both neighbours (or strictly smaller than both), you have found a critical point at that index. Now the problem stops being about linked lists at all: it becomes "given a stream of increasing indices, find the smallest and largest gap."

### Hint 3

You do not need to store the indices in a slice. For the **maximum** distance, the answer is always `lastCriticalIndex - firstCriticalIndex`, because the extremes are the farthest apart pair by definition. For the **minimum** distance, the closest pair must be *adjacent* in sorted order — any non-adjacent pair straddles an adjacent pair and is therefore at least as wide. So remember only three numbers as you go: the first critical index, the previous critical index, and the running minimum gap. That collapses the space to O(1).

## Approach

The two halves of the answer come from the same single pass, but they use different bookkeeping.

**Step 1 — detect critical points.** A node is critical only if it has both a predecessor and a successor, so the head and the tail are never candidates. Traverse with a sliding triple `(prev, cur, next)` starting at `prev = head`, `cur = head.Next`, and keep a 1-based index for `cur`. At each step:

- local maxima: `cur.Val > prev.Val && cur.Val > next.Val`
- local minima: `cur.Val < prev.Val && cur.Val < next.Val`

Note that both comparisons are *strict*. A plateau such as `[2,2,2]` produces no critical point, which is exactly why example 3 has only two critical points despite the list wobbling several times.

**Step 2 — track distances in O(1) space.** Maintain three variables:

- `first` — index of the first critical point seen (set once).
- `prev` — index of the most recently seen critical point.
- `minDist` — the smallest gap seen so far.

When a new critical point at index `i` appears: if `first` is unset, record `first = i`. Otherwise update `minDist = min(minDist, i - prev)`. In both cases set `prev = i`.

Why is looking only at adjacent gaps enough for the minimum? The critical indices are discovered in increasing order. For any pair `i < j` that are not adjacent in that order, there is some `k` with `i < k < j`, and `k - i < j - i`. So a non-adjacent pair can never be the unique closest pair — checking consecutive gaps covers every candidate. Symmetrically, the maximum is pinned by the two extremes, so after the pass `maxDist = prev - first` (where `prev` now holds the last critical index).

**Step 3 — the fewer-than-two case.** If we ended with zero or one critical point, no pair exists and the answer is `[-1, -1]`. Detect this by checking whether `minDist` was ever updated, or equivalently whether `first == prev`.

**Worked example:** `head = [5,3,1,2,5,1,2]`, 1-based indices.

| index | prev, cur, next | critical? | first | prev crit | minDist |
|-------|-----------------|-----------|-------|-----------|---------|
| 2 | 5, 3, 1 | no (3 is between) | – | – | ∞ |
| 3 | 3, 1, 2 | minima | 3 | 3 | ∞ |
| 4 | 1, 2, 5 | no | 3 | 3 | ∞ |
| 5 | 2, 5, 1 | maxima | 3 | 5 | 2 |
| 6 | 5, 1, 2 | minima | 3 | 6 | 1 |

The loop stops at index 6 because index 7 is the tail. Result: `minDist = 1`, `maxDist = 6 - 3 = 3`, i.e. `[1, 3]` — matching the expected output.

This is an approachable Medium: the detection rule is simple, and the real learning is the "closest pair must be adjacent" observation that lets you drop the array of indices entirely. The two traps that actually cost people submissions are using non-strict comparisons (breaking on plateaus) and forgetting the fewer-than-two guard.

## Complexity Analysis

Time Complexity: O(n) — one pass, visiting each node a constant number of times.
Space Complexity: O(1) — only a handful of integers and node pointers, regardless of list length.

## Edge Cases

- **Fewer than two critical points** (`[3,1]`, `[2,1,3]`) — must return `[-1, -1]`. A list with exactly one critical point is the sneaky version of this: the detection code runs and succeeds, but there is still no pair, so the guard must be on the *count*, not on whether any critical point exists.
- **Minimum-length list** (`n == 2`, guaranteed by the constraints as the floor) — no node has both neighbours, so the traversal body never executes. The code must not dereference `head.Next.Next`.
- **Plateaus / equal adjacent values** (`[1,3,2,2,3,2,2,2,7]`, `[5,5,5,5,5]`) — because both comparisons are strict, flat runs contribute nothing. Using `>=` or `<=` here silently produces wrong answers on the given examples.
- **Head and tail are never critical** — example 3 explicitly calls this out: the final `7` is larger than its predecessor but has no successor, so it does not count.
- **Strictly alternating list** (`[1,5,1,5,1,5,1]`) — nearly every interior node is critical, so `minDist` is 1 and `maxDist` spans almost the whole list. Good check that you are comparing consecutive gaps rather than only the first pair.
- **Monotonic list** (`[1,2,3,4,5]`) — zero critical points; exercises the `[-1, -1]` path from the other direction.
- **Overflow / initialisation** — initialise `minDist` to a sentinel larger than any possible gap (`n` itself, or `math.MaxInt32`) rather than `0`, otherwise the `min` update can never fire.
