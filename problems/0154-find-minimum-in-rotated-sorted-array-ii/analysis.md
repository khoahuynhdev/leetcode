# 0154. Find Minimum in Rotated Sorted Array II

[LeetCode Link](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/)

Difficulty: Hard
Topics: Array, Binary Search
Acceptance Rate: 45.2%

## Hints

### Hint 1

The array is "sorted then rotated", which means even though it is no longer fully sorted, it still has structure: it consists of two non-decreasing halves with one drop in between. Whenever you see a sorted-ish array and the constraint to be efficient, think about binary search rather than a linear scan.

### Hint 2

In the duplicate-free version (problem 153), you can always compare `nums[mid]` with one of the endpoints to decide which half contains the minimum. Compare `nums[mid]` with `nums[right]`:
- If `nums[mid] > nums[right]`, the drop is to the right of `mid`, so the minimum lies in `(mid, right]`.
- If `nums[mid] < nums[right]`, the right half is sorted, so the minimum lies in `[left, mid]`.

What goes wrong with this comparison when duplicates are allowed?

### Hint 3

When `nums[mid] == nums[right]`, you genuinely cannot tell which side the rotation point lives on. For example, both `[1,1,1,0,1]` and `[1,0,1,1,1]` give `nums[mid] == nums[right]`, yet the minimum sits on opposite sides. The safe move is to simply shrink the search window by one — `right--` — and continue. This still discards `nums[right]` from consideration because we know `nums[mid]` (which is equal to it) is still a candidate, so we lose nothing.

## Approach

Use binary search over the index range `[left, right]` and shrink it based on how `nums[mid]` compares to `nums[right]`:

1. Initialize `left = 0`, `right = len(nums) - 1`.
2. While `left < right`:
   - `mid = left + (right - left) / 2`.
   - If `nums[mid] > nums[right]`: the minimum must be strictly to the right of `mid`. Set `left = mid + 1`.
   - Else if `nums[mid] < nums[right]`: the segment `(mid, right]` is sorted, so the minimum is in `[left, mid]`. Set `right = mid`.
   - Else (`nums[mid] == nums[right]`): we cannot decide. Safely shrink by `right--`. We do not lose the minimum because `nums[mid]` is still in the search window and equals `nums[right]`.
3. Return `nums[left]`.

**Walkthrough on `[2,2,2,0,1]`**:
- `left=0, right=4, mid=2`: `nums[2]=2 > nums[4]=1` → `left=3`.
- `left=3, right=4, mid=3`: `nums[3]=0 < nums[4]=1` → `right=3`.
- Loop exits, return `nums[3] = 0`.

**Walkthrough on `[10,1,10,10,10]`** (showing the duplicate case):
- `left=0, right=4, mid=2`: `nums[2]=10 == nums[4]=10` → `right=3`.
- `left=0, right=3, mid=1`: `nums[1]=1 < nums[3]=10` → `right=1`.
- `left=0, right=1, mid=0`: `nums[0]=10 > nums[1]=1` → `left=1`.
- Loop exits, return `nums[1] = 1`.

Why compare against `nums[right]` rather than `nums[left]`? Comparing with the right endpoint cleanly separates the "rotation lies to the right" case (`nums[mid] > nums[right]`) from the "we are already in the right-sorted run" case (`nums[mid] < nums[right]`). Comparing with `nums[left]` is also possible but requires an extra case to handle the not-actually-rotated input.

## Complexity Analysis

Time Complexity: O(log n) average, O(n) worst case. The `right--` branch only fires when `nums[mid] == nums[right]`, and in the pathological case where all elements are equal (e.g., `[3,3,3,3,3]`) every iteration takes that branch, degrading the algorithm to linear time.

Space Complexity: O(1). Only a constant number of index variables are used.

## Edge Cases

- **Single element**: `[1]` — loop never executes, returns `nums[0]`.
- **Not actually rotated** (rotated `n` times): e.g., `[1,2,3,4,5]` — `nums[mid] < nums[right]` always holds, so `right` shrinks until it equals `left = 0`, returning the first element.
- **All elements equal**: `[3,3,3,3]` — the `nums[mid] == nums[right]` branch fires every iteration, giving the O(n) worst case but still correct.
- **Duplicates straddling the rotation point**: `[10,1,10,10,10]` and `[1,1,1,0,1]` — the case the duplicates rule was designed for; relies on `right--` to safely make progress.
- **Two elements**: `[2,1]` or `[1,2]` — one iteration suffices; verifies the `left < right` termination is correct.
- **Negative numbers**: values can be negative; nothing in the algorithm assumes positivity.
