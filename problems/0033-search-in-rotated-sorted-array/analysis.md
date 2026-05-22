# 0033. Search in Rotated Sorted Array

[LeetCode Link](https://leetcode.com/problems/search-in-rotated-sorted-array/)

Difficulty: Medium
Topics: Array, Binary Search
Acceptance Rate: 44.6%

## Hints

### Hint 1

The required `O(log n)` runtime is a strong signal. On a plain sorted array
that complexity points straight at **binary search**. The twist here is that the
array has been rotated, so it is not globally sorted. Ask yourself: can binary
search still work if only *part* of the array is sorted at any given moment?

### Hint 2

At every step, pick a middle index and compare the values at the boundaries with
the value at `mid`. A rotated array always splits into **one sorted half and one
that contains the rotation point**. If you can reliably tell which half is the
clean, sorted one, you can decide whether the target lives inside that sorted
range — and discard half the array each iteration, just like normal binary
search.

### Hint 3

The key test is `nums[lo] <= nums[mid]`.

- If it holds, the **left half `[lo..mid]` is sorted**. The target belongs there
  only if `nums[lo] <= target < nums[mid]`; otherwise search the right half.
- If it fails, the **right half `[mid..hi]` is sorted**. The target belongs there
  only if `nums[mid] < target <= nums[hi]`; otherwise search the left half.

Use `<=` (not `<`) in the sorted-half check so that a two-element window where
`lo == mid` is still classified correctly.

## Approach

We run a single modified binary search. Maintain a window `[lo, hi]` and loop
while `lo <= hi`.

1. Compute `mid = lo + (hi-lo)/2` (the `lo + ...` form avoids integer overflow).
2. If `nums[mid] == target`, we are done — return `mid`.
3. Otherwise figure out which half is sorted:
   - **Left half sorted** (`nums[lo] <= nums[mid]`): if the target falls within
     `[nums[lo], nums[mid])`, narrow to the left (`hi = mid - 1`); else go right
     (`lo = mid + 1`).
   - **Right half sorted** (else): if the target falls within
     `(nums[mid], nums[hi]]`, narrow to the right (`lo = mid + 1`); else go left
     (`hi = mid - 1`).
4. If the loop ends without a hit, the target is absent — return `-1`.

Why it works: a rotated sorted array of distinct values is two sorted runs glued
together. For any `mid`, at least one of the two halves `[lo..mid]` or
`[mid..hi]` is fully sorted (it cannot contain the rotation seam). Once we know a
half is sorted, a simple range check tells us definitively whether `target` could
be inside it. We then keep the half that might contain the target and throw away
the other, halving the search space every iteration.

Walkthrough on `nums = [4,5,6,7,0,1,2]`, `target = 0`:

- `lo=0, hi=6, mid=3` → `nums[3]=7`. Left half `[4..7]` is sorted; `0` is not in
  `[4,7)`, so go right: `lo=4`.
- `lo=4, hi=6, mid=5` → `nums[5]=1`. Left half `[0,1]` is sorted; `0` is in
  `[0,1)`, so go left: `hi=4`.
- `lo=4, hi=4, mid=4` → `nums[4]=0` equals target. Return `4`.

## Complexity Analysis

Time Complexity: O(log n) — each iteration discards half of the remaining window.
Space Complexity: O(1) — only a few index variables, no extra allocation.

## Edge Cases

- **Single element** (`[1]`): when `target` is absent (e.g. `0`) the loop runs
  once and returns `-1`; when it matches, it returns `0`. The `lo + (hi-lo)/2`
  midpoint and `lo <= hi` guard keep this correct.
- **Empty array**: `hi = len-1 = -1`, so the loop body never executes and we
  return `-1` immediately. (The stated constraints say `nums.length >= 1`, but
  guarding for it costs nothing.)
- **Not actually rotated** (`[1,2,3,4,5]`): `nums[lo] <= nums[mid]` is always
  true, so the algorithm degrades gracefully into ordinary binary search.
- **Target at the boundaries** (first or last element, or adjacent to the
  rotation seam): the inclusive/exclusive bounds in the range checks
  (`nums[lo] <= target < nums[mid]` and `nums[mid] < target <= nums[hi]`) are
  what keep these from being missed — getting one `<` vs `<=` wrong is the most
  common bug.

This is a deservedly classic Medium: the binary-search skeleton is familiar, but
the "which half is sorted, and is the target in it" reasoning takes a couple of
careful passes to internalize. Once the boundary conditions click, it is a clean
and very reusable pattern.
