# 1846. Maximum Element After Decreasing and Rearranging

[LeetCode Link](https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/)

Difficulty: Medium
Topics: Array, Greedy, Sorting

Acceptance Rate: 66.7%

## Hints

### Hint 1

The two operations you are given are very generous: you can rearrange the
elements in any order, and you can lower any element to any smaller positive
integer. Because order is fully under your control, the *positions* of the
values don't really matter — only the *multiset* of values does. Ask yourself:
if you could freely sort the array, what arrangement gives you the best shot at
a tall final value?

### Hint 2

Once you sort the array in non-decreasing order, the constraints become easy to
read. The first element must be `1`, and each subsequent element may be at most
`1` larger than the one before it. Notice that decreasing is allowed but
increasing is not — so the value at each position can never "jump up" to fill a
gap. Think greedily, left to right.

### Hint 3

After sorting, walk from left to right keeping the value you have committed to
the previous position. Start that value at `1`. For the current sorted element,
the best you can legally keep is `min(arr[i], prev + 1)` — you can always
*decrease* down to `prev + 1`, but you can never raise an element above its
original value, and you can never exceed `prev + 1`. The answer is the value
committed to the very last position, which (because the array is sorted) is the
overall maximum.

## Approach

The key realization is that rearranging is free, so we may as well sort `arr` in
non-decreasing order. In a sorted array the optimal strategy is to make the
sequence climb as fast as the rules allow.

Algorithm:

1. Sort `arr` ascending.
2. Force the first element to `1` (we can always decrease it; it must be `1`).
3. Iterate `i` from `1` to `n - 1`, tracking `prev`, the value we committed to
   position `i - 1` (initially `1`).
   - The value we keep at position `i` is `min(arr[i], prev + 1)`:
     - We may decrease `arr[i]` down to `prev + 1` to respect the adjacency rule.
     - We may **not** increase `arr[i]`, so if `arr[i]` is already `<= prev + 1`
       we just keep `arr[i]` (which keeps the array valid and non-decreasing).
   - Set `prev` to that committed value.
4. The answer is `prev` after the loop — the value at the last (largest) sorted
   position.

Why it works: sorting guarantees that at each step `arr[i] >= arr[i-1]`, so the
committed values are non-decreasing and the maximum lands on the last position.
Greedily taking `min(arr[i], prev + 1)` is optimal because committing any larger
value would violate a rule, and committing any smaller value would needlessly
cap every later position (each later value is bounded by `prev + 1`, so a
smaller `prev` can only hurt).

Worked example: `arr = [100, 1, 1000]`.

- Sort → `[1, 100, 1000]`.
- Position 0: force to `1`. `prev = 1`.
- Position 1: `min(100, 1 + 1) = 2`. `prev = 2`.
- Position 2: `min(1000, 2 + 1) = 3`. `prev = 3`.
- Answer: `3`.

## Complexity Analysis

Time Complexity: O(n log n) — dominated by sorting; the linear scan is O(n).
Space Complexity: O(1) extra (or O(log n) for the sort's recursion stack),
ignoring the in-place sort of the input.

## Edge Cases

- **Single element** (`len(arr) == 1`): The only element must become `1`, so the
  answer is always `1` regardless of its original value.
- **First element already greater than 1**: We must decrease it to `1`; the loop
  starts `prev` at `1` to enforce this.
- **All equal values** (e.g. `[5, 5, 5]`): After sorting they still climb by at
  most `1` each step, yielding `1, 2, 3, ...`.
- **Already valid sequence** (e.g. `[1, 2, 3, 4, 5]`): No decreases are needed;
  the answer is simply the maximum, which `min(arr[i], prev + 1)` preserves.
- **Large gaps / large values** (up to `1e9`): The committed value can grow by at
  most `1` per element, so it never exceeds `n`, keeping everything well within
  `int` range.
