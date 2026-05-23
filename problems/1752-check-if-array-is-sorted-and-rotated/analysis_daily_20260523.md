# 1752. Check if Array Is Sorted and Rotated

[LeetCode Link](https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/)

Difficulty: Easy
Topics: Array
Acceptance Rate: 56.2%

## Hints

### Hint 1

Don't try to "un-rotate" the array or guess the original sorted version. Instead, think about what property a sorted-and-rotated array has when you read it left to right, treating it as if it wraps around (the last element connects back to the first).

### Hint 2

In a non-decreasing array, every element is `<=` the next one. Rotating that array breaks this rule in exactly one place: the single spot where a larger value is immediately followed by a smaller one (the "wrap-around seam"). Try counting how many such drops exist as you scan the array circularly.

### Hint 3

Walk the array once and count positions `i` where `nums[i] > nums[(i+1) % n]`. A non-rotated sorted array has **0** such drops; a properly rotated one has exactly **1**. If you find **2 or more**, it can never be a rotation of a sorted array, so return `false`. The answer is simply: `count <= 1`.

## Approach

The key realization is that a rotation of a non-decreasing array introduces at most one "break point" — the place where the sequence drops from a high value back down to the smallest value.

Concretely, imagine the original sorted array `[1,2,3,4,5]`. Every adjacent pair satisfies `nums[i] <= nums[i+1]`. When we rotate it to `[3,4,5,1,2]`, all pairs still satisfy that rule except one: `5 -> 1`. That single drop is the seam created by the rotation.

So the algorithm is:

1. Let `n` be the length of the array.
2. Scan each index `i` from `0` to `n-1` and compare `nums[i]` with its circular neighbor `nums[(i+1) % n]`. The modulo lets the last element compare against the first, closing the loop.
3. Count how many times `nums[i] > nums[(i+1) % n]` (a "drop").
4. If the count is `0` (already sorted, rotation by zero) or `1` (one valid seam), return `true`. If it's `2` or more, return `false`.

Why does counting the circular comparison work so cleanly? Because a valid sorted-and-rotated array is just a non-decreasing sequence wrapped into a circle. A circle of non-decreasing values can have at most one place where it "wraps" from the maximum back to the minimum. More than one such wrap means there's no way to cut the circle into a single non-decreasing line.

Example trace on `[3,4,5,1,2]` (`n = 5`):

- `3 <= 4` ok
- `4 <= 5` ok
- `5 > 1` drop (count = 1)
- `1 <= 2` ok
- `2 <= 3` (wrap: index 4 vs index 0) ok

Total drops = 1, so the result is `true`.

On `[2,1,3,4]` (`n = 4`):

- `2 > 1` drop (count = 1)
- `1 <= 3` ok
- `3 <= 4` ok
- `4 > 2` (wrap) drop (count = 2)

Total drops = 2, so the result is `false`.

## Complexity Analysis

Time Complexity: O(n) — a single pass over the array.
Space Complexity: O(1) — only a counter is used.

## Edge Cases

- **Single-element array** (`[1]`): A length-1 array is trivially sorted; the circular comparison compares the element with itself (`nums[0] > nums[0]` is false), giving 0 drops. Returns `true`.
- **Already sorted, no rotation** (`[1,2,3]`): Rotation by zero is allowed, so 0 drops must return `true`.
- **Duplicates** (`[1,1,1]` or `[2,1,2,2]`): The problem allows duplicates in the original array. Using strict `>` for the drop comparison (not `>=`) is essential — equal adjacent values are valid in a non-decreasing array and must not be counted as drops.
- **The wrap-around comparison itself**: Forgetting to compare the last element against the first (`nums[n-1]` vs `nums[0]`) would miss the seam for fully-sorted arrays that look fine internally but aren't a valid rotation; including it via `(i+1) % n` handles this correctly.
- **Strictly decreasing array** (`[3,2,1]`): Produces 2 drops (`3>2`, `2>1`), correctly returning `false`.
