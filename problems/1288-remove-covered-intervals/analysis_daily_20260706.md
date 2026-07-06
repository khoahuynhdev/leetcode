# 1288. Remove Covered Intervals

[LeetCode Link](https://leetcode.com/problems/remove-covered-intervals/)

Difficulty: Medium
Topics: Array, Sorting
Acceptance Rate: 57.0%

## Hints

### Hint 1

An interval is "covered" only in relation to *another* interval. Comparing every
pair is O(n²) and workable given the small constraints, but there is a cleaner
angle. When a problem asks you to reason about which intervals contain which,
think about what happens if you first put them in a predictable order. Which
sorting key would let you decide "covered or not" by looking only at what you've
already seen?

### Hint 2

Sort the intervals by their **start** value in ascending order. After sorting, if
interval `A` comes before interval `B`, then `A.start <= B.start`. That already
satisfies the first half of the coverage condition (`c <= a`). Now coverage
reduces to a single comparison on the **end** values. As you sweep left to right,
what single number do you need to remember to know whether the current interval is
swallowed by something earlier?

### Hint 3

Track the maximum end value seen so far. As you walk the sorted list, the current
interval is covered when its end is `<= maxEnd`; otherwise it survives and you
extend `maxEnd`. The one trap: two intervals sharing the same start. If `[1,4]`
and `[1,10]` are compared, `[1,4]` is covered by `[1,10]`, but a naive start-only
sort might place `[1,4]` first and wrongly count it. Fix this with the tie-break:
when starts are equal, sort by end in **descending** order so the longer interval
is processed first and correctly "eats" the shorter one.

## Approach

Sorting turns an all-pairs comparison into a single linear sweep.

1. **Sort** the intervals by start ascending. On ties (equal starts), sort by end
   descending. This guarantees that for any interval we look at, every previously
   seen interval has a start less than or equal to the current start, and among
   equal starts the widest interval is seen first.

2. **Sweep** through the sorted list, keeping `maxEnd`, the largest end value among
   all intervals processed so far. For each interval `[l, r]`:
   - If `r <= maxEnd`, then some earlier interval starts no later (guaranteed by
     the sort) and ends no earlier, so `[l, r]` is covered — skip it.
   - Otherwise `[l, r]` is not covered; increment the surviving count and update
     `maxEnd = r`.

3. Return the count of surviving intervals.

**Why the tie-break matters:** Consider `[[1,4],[1,10]]`. Both start at 1. If we
processed `[1,4]` first, `maxEnd` becomes 4, then `[1,10]` has end `10 > 4` and is
counted — giving 2. But `[1,4]` is actually covered by `[1,10]`, so the answer is
1. Sorting equal starts by end descending processes `[1,10]` first (`maxEnd = 10`),
then `[1,4]` has `4 <= 10` and is correctly discarded.

**Walkthrough on Example 1** `[[1,4],[3,6],[2,8]]`:
- Sorted (start asc, end desc): `[[1,4],[2,8],[3,6]]`.
- `[1,4]`: `4 > 0`, survives. count = 1, maxEnd = 4.
- `[2,8]`: `8 > 4`, survives. count = 2, maxEnd = 8.
- `[3,6]`: `6 <= 8`, covered. count stays 2.
- Result: **2**. ✅

## Complexity Analysis

Time Complexity: O(n log n) — dominated by the sort; the sweep is O(n).
Space Complexity: O(1) extra (or O(log n)–O(n) depending on the sort's internal
allocation), since we sort in place and keep only a couple of scalars.

## Edge Cases

- **Single interval:** `[[a, b]]` can never be covered by another — the answer is
  always 1. The sweep handles this naturally.
- **Equal starts:** e.g. `[[1,4],[1,10]]`. Without the end-descending tie-break the
  shorter interval is miscounted. This is the single most important case to get
  right.
- **Identical/nested chains:** e.g. `[[1,2],[1,4],[1,6],[1,8]]` all share a start;
  only the widest survives, so the answer is 1.
- **No coverage at all:** disjoint or partially overlapping intervals like
  `[[1,2],[3,4],[5,6]]` — none is covered, so every interval counts.
- **Interval that shares an end but has a later start:** e.g. `[[2,8],[3,8]]`;
  `[3,8]` is covered because `2 <= 3` and `8 <= 8`. The `<=` comparison (not `<`)
  is essential here.
