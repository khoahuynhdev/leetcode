# 2089. Find Target Indices After Sorting Array

[LeetCode Link](https://leetcode.com/problems/find-target-indices-after-sorting-array/)

Difficulty: Easy
Topics: Array, Binary Search, Sorting
Acceptance Rate: 79.4%

## Hints

### Hint 1

The obvious route is to sort and then scan, and that works fine. But ask yourself: do you actually need the whole sorted array, or only to know *where* the target values end up in it?

### Hint 2

In a sorted array, all copies of `target` sit next to each other in one contiguous block. If you know where that block starts and how long it is, you know the answer. What determines where the block starts?

### Hint 3

The block starts at index `less`, where `less` is the number of elements strictly smaller than `target`. Its length is `equal`, the number of elements equal to `target`. One pass counting both gives the answer `[less, less+1, ..., less+equal-1]`, with no sorting needed.

## Approach

**Counting (optimal):**

1. Walk through `nums` once, keeping two counters:
   - `less`: how many elements are `< target`
   - `equal`: how many elements are `== target`
2. After sorting, every smaller element comes before every copy of `target`, so the first copy of `target` sits at index `less`.
3. The copies take up `equal` consecutive slots, so the answer is `less, less+1, ..., less+equal-1`.
4. If `equal == 0`, the result is an empty list.

**Example:** `nums = [1,2,5,2,3]`, `target = 2`
- Elements < 2: `[1]`, so `less = 1`
- Elements == 2: `[2, 2]`, so `equal = 2`
- Answer: `[1, 2]`, which matches the sorted array `[1,2,2,3,5]`.

**Alternative approaches:**
- *Sort + scan:* sort, then collect every index where `nums[i] == target`. O(n log n).
- *Sort + binary search:* sort, then use lower bound and upper bound to find the block. Still O(n log n) because of the sort, but it's good practice for the binary search topic tag.

The counting approach is better than both because it never builds the sorted array. It only figures out where the target's block would land.

## Complexity Analysis

Time Complexity: O(n), one pass to count, plus output of at most n indices.
Space Complexity: O(1) extra (not counting the output slice).

## Edge Cases

- **Target not present:** `equal == 0`, so return an empty (non-nil) slice instead of `nil`, to match the expected `[]`.
- **All elements equal target:** `less == 0`, `equal == n`, so the answer is every index `0..n-1`.
- **Target smaller than every element:** nothing is less and nothing is equal, so the result is empty.
- **Target larger than every element:** `less == n`, `equal == 0`, so the result is empty. Make sure you don't produce out-of-range indices.
- **Single-element array:** either `[0]` or `[]`, depending on whether it matches.
